package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/image_database/internal/config"
	"github.com/ManyakRus/image_database/internal/types"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TableColumn struct {
	TableName          string `json:"table_name"   gorm:"column:table_name;default:''"`
	TableComment       string `json:"table_comment"   gorm:"column:table_comment;default:''"`
	ColumnName         string `json:"column_name"   gorm:"column:column_name;default:''"`
	ColumnType         string `json:"type_name"   gorm:"column:type_name;default:''"`
	ColumnIsPrimaryKey string `json:"is_primary_key"   gorm:"column:is_primary_key;default:''"`
	ColumnDescription  string `json:"description"   gorm:"column:description;default:''"`
	ColumnTableKey     string `json:"table_key"   gorm:"column:table_key;default:''"`
	ColumnColumnKey    string `json:"column_key"   gorm:"column:column_key;default:''"`
}

// FillMapTable - возвращает массив MassTable данными из БД
func FillMapTable() (map[string]*types.Table, error) {
	var err error
	//MassTable := make([]types.Table, 0)
	MapTable := make(map[string]*types.Table, 0)

	TextSQL := `

drop table if exists temp_keys; 
CREATE TEMPORARY TABLE temp_keys (table_from text,  column_from text, table_to text, column_to text);

------------------------------------------- Все внешние ключи ------------------------------
insert into temp_keys
SELECT 
       (select r.relname from pg_class r where r.oid = c.conrelid) as table_from,
       UNNEST((select array_agg(attname) from pg_attribute where attrelid = c.conrelid and array[attnum] <@ c.conkey)) as column_from,
       (select  r.relname from pg_class r where r.oid = c.confrelid) as table_to,
       a.attname as column_to
FROM 
	pg_constraint c 
	
join 
	pg_attribute a 
on 
	c.confrelid=a.attrelid and a.attnum = ANY(confkey)
	
WHERE 1=1
	--and c.confrelid = (select oid from pg_class where relname = 'debt_types')
	AND c.confrelid!=c.conrelid
;

------------------------------------------- Все primary keys ------------------------------
drop table if exists temp_primary_keys; 
CREATE TEMPORARY TABLE temp_primary_keys (table_name text,  column_name text);

insert into temp_primary_keys
select 
    ccu.table_name,
	(ccu.column_name) as column_name
       
from pg_constraint pgc
         join pg_namespace nsp on nsp.oid = pgc.connamespace
         join pg_class  cls on pgc.conrelid = cls.oid
         left join information_schema.constraint_column_usage ccu
                   on pgc.conname = ccu.constraint_name
                       and nsp.nspname = ccu.constraint_schema
WHERE 1=1
	and ccu.table_schema = 'public'
	and contype = 'p'
	
--GROUP BY
--	ccu.table_name
--HAVING sum(1)=1
;


------------------------------------------- Все таблицы и колонки ------------------------------

SELECT 
	c.table_name, 
	c.column_name,
	c.udt_name as type_name,
    (SELECT obj_description(oid) FROM pg_class as r WHERE relkind = 'r' and r.oid = st.relid) as table_comment,
	CASE
		WHEN tpk.table_name is not null
		THEN true
		ELSE false END 
	    as is_primary_key,
	COALESCE(pgd.description, '') as description,
	COALESCE(keys.table_to, '') as table_key,
	COALESCE(keys.column_to, '') as column_key 
	
FROM 
	information_schema.columns c 
	
left join 
	pg_catalog.pg_statio_all_tables as st
on 
	c.table_schema = st.schemaname
	and c.table_name   = st.relname

	    
left join 
	pg_catalog.pg_description pgd 
on 
	pgd.objoid = st.relid
	and pgd.objsubid   = c.ordinal_position


LEFT JOIN --внешние ключи
	temp_keys as keys
ON
	keys.table_from = c.table_name
	and keys.column_from = c.column_name

	
LEFT JOIN --вьюхи
	INFORMATION_SCHEMA.views as v
ON
	v.table_schema = 'public'
	and v.table_name = c.table_name


LEFT JOIN
	temp_primary_keys as tpk
ON 
	tpk.table_name = c.table_name
	and tpk.column_name = c.column_name


where 1=1
	and c.table_schema='public'
	and v.table_name is null
	--INCLUDE_TABLES
	--EXCLUDE_TABLES

order by 
	table_name, 
	is_primary_key desc,
	column_name
`

	SCHEMA := strings.Trim(postgres_gorm.Settings.DB_SCHEMA, " ")
	if SCHEMA != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "public", SCHEMA)
	}

	if config.Settings.INCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--INCLUDE_TABLES", "and c.table_name ~* '"+config.Settings.INCLUDE_TABLES+"'")
	}

	if config.Settings.EXCLUDE_TABLES != "" {
		TextSQL = strings.ReplaceAll(TextSQL, "--EXCLUDE_TABLES", "and c.table_name !~* '"+config.Settings.EXCLUDE_TABLES+"'")
	}

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*60)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	db.WithContext(ctx)

	//запрос
	//запустим все запросы отдельно
	var tx *gorm.DB
	sqlSlice := strings.Split(TextSQL, ";")
	len1 := len(sqlSlice)
	for i, TextSQL1 := range sqlSlice {
		//batch.Queue(TextSQL1)
		if i == len1-1 {
			tx = db.Raw(TextSQL1)
		} else {
			tx = db.Exec(TextSQL1)
			//rows.Close()
		}
		err = tx.Error
		if err != nil {
			log.Panic("DB.Raw() error:", err)
		}
	}

	//tx := db.Raw(TextSQL)
	//err = tx.Error
	//if err != nil {
	//	sError := fmt.Sprint("db.Raw() error: ", err)
	//	log.Panicln(sError)
	//	return MassTable, err
	//}

	//ответ в структуру
	MassTableColumn := make([]TableColumn, 0)
	tx = tx.Scan(&MassTableColumn)
	err = tx.Error
	if err != nil {
		sError := fmt.Sprint("Get_error()  error: ", err)
		log.Panicln(sError)
		return MapTable, err
	}

	//проверка 0 строк
	if tx.RowsAffected == 0 {
		sError := fmt.Sprint("db.Raw() RowsAffected =0 ")
		log.Warn(sError)
		err = errors.New(sError)
		//log.Panicln(sError)
		return MapTable, err
	}

	//заполним MapTable
	MapColumns := make(map[string]types.Column, 0)
	OrderNumberColumn := 0
	OrderNumberTable := 0
	TableName0 := ""
	Table1 := CreateTable()
	for _, v := range MassTableColumn {
		if v.TableName != TableName0 {
			OrderNumberColumn = 0
			Table1.MapColumns = MapColumns
			MapColumns = make(map[string]types.Column, 0)
			if TableName0 != "" {
				//MassTable = append(MassTable, Table1)
				MapTable[TableName0] = Table1
				OrderNumberTable++
			}
			Table1 = CreateTable()
			Table1.Name = v.TableName
			Table1.Comment = v.TableComment
			Table1.OrderNumber = OrderNumberTable
		}

		Column1 := types.Column{}
		Column1.Name = v.ColumnName
		Column1.Type = v.ColumnType
		if v.ColumnIsPrimaryKey == "true" {
			Column1.IsPrimaryKey = true
		}
		Column1.Description = v.ColumnDescription
		Column1.OrderNumber = OrderNumberColumn
		Column1.TableKey = v.ColumnTableKey
		Column1.ColumnKey = v.ColumnColumnKey

		MapColumns[v.ColumnName] = Column1
		//Table1.Columns = append(Table1.Columns, Column1)

		OrderNumberColumn++
		TableName0 = v.TableName
	}
	if Table1.Name != "" {
		Table1.MapColumns = MapColumns
		MapTable[TableName0] = Table1
	}

	return MapTable, err
}

func CreateTable() *types.Table {
	Otvet := &types.Table{}
	Otvet.MapColumns = make(map[string]types.Column, 0)

	return Otvet
}
