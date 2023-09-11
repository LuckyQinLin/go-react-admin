package gorm

import (
	"admin-api/internal/gorm/clause"
	"admin-api/internal/gorm/logger"
	"admin-api/internal/gorm/schema"
	"context"
	"database/sql"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestConfig_AfterInitialize(t *testing.T) {
	type fields struct {
		SkipDefaultTransaction                   bool
		NamingStrategy                           schema.Namer
		FullSaveAssociations                     bool
		Logger                                   logger.Interface
		NowFunc                                  func() time.Time
		DryRun                                   bool
		PrepareStmt                              bool
		DisableAutomaticPing                     bool
		DisableForeignKeyConstraintWhenMigrating bool
		IgnoreRelationshipsWhenMigrating         bool
		DisableNestedTransaction                 bool
		AllowGlobalUpdate                        bool
		QueryFields                              bool
		CreateBatchSize                          int
		TranslateError                           bool
		ClauseBuilders                           map[string]clause.ClauseBuilder
		ConnPool                                 ConnPool
		Dialector                                Dialector
		Plugins                                  map[string]Plugin
		callbacks                                *callbacks
		cacheStore                               *sync.Map
	}
	type args struct {
		db *DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				SkipDefaultTransaction:                   tt.fields.SkipDefaultTransaction,
				NamingStrategy:                           tt.fields.NamingStrategy,
				FullSaveAssociations:                     tt.fields.FullSaveAssociations,
				Logger:                                   tt.fields.Logger,
				NowFunc:                                  tt.fields.NowFunc,
				DryRun:                                   tt.fields.DryRun,
				PrepareStmt:                              tt.fields.PrepareStmt,
				DisableAutomaticPing:                     tt.fields.DisableAutomaticPing,
				DisableForeignKeyConstraintWhenMigrating: tt.fields.DisableForeignKeyConstraintWhenMigrating,
				IgnoreRelationshipsWhenMigrating:         tt.fields.IgnoreRelationshipsWhenMigrating,
				DisableNestedTransaction:                 tt.fields.DisableNestedTransaction,
				AllowGlobalUpdate:                        tt.fields.AllowGlobalUpdate,
				QueryFields:                              tt.fields.QueryFields,
				CreateBatchSize:                          tt.fields.CreateBatchSize,
				TranslateError:                           tt.fields.TranslateError,
				ClauseBuilders:                           tt.fields.ClauseBuilders,
				ConnPool:                                 tt.fields.ConnPool,
				Dialector:                                tt.fields.Dialector,
				Plugins:                                  tt.fields.Plugins,
				callbacks:                                tt.fields.callbacks,
				cacheStore:                               tt.fields.cacheStore,
			}
			if err := c.AfterInitialize(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("AfterInitialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_Apply(t *testing.T) {
	type fields struct {
		SkipDefaultTransaction                   bool
		NamingStrategy                           schema.Namer
		FullSaveAssociations                     bool
		Logger                                   logger.Interface
		NowFunc                                  func() time.Time
		DryRun                                   bool
		PrepareStmt                              bool
		DisableAutomaticPing                     bool
		DisableForeignKeyConstraintWhenMigrating bool
		IgnoreRelationshipsWhenMigrating         bool
		DisableNestedTransaction                 bool
		AllowGlobalUpdate                        bool
		QueryFields                              bool
		CreateBatchSize                          int
		TranslateError                           bool
		ClauseBuilders                           map[string]clause.ClauseBuilder
		ConnPool                                 ConnPool
		Dialector                                Dialector
		Plugins                                  map[string]Plugin
		callbacks                                *callbacks
		cacheStore                               *sync.Map
	}
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				SkipDefaultTransaction:                   tt.fields.SkipDefaultTransaction,
				NamingStrategy:                           tt.fields.NamingStrategy,
				FullSaveAssociations:                     tt.fields.FullSaveAssociations,
				Logger:                                   tt.fields.Logger,
				NowFunc:                                  tt.fields.NowFunc,
				DryRun:                                   tt.fields.DryRun,
				PrepareStmt:                              tt.fields.PrepareStmt,
				DisableAutomaticPing:                     tt.fields.DisableAutomaticPing,
				DisableForeignKeyConstraintWhenMigrating: tt.fields.DisableForeignKeyConstraintWhenMigrating,
				IgnoreRelationshipsWhenMigrating:         tt.fields.IgnoreRelationshipsWhenMigrating,
				DisableNestedTransaction:                 tt.fields.DisableNestedTransaction,
				AllowGlobalUpdate:                        tt.fields.AllowGlobalUpdate,
				QueryFields:                              tt.fields.QueryFields,
				CreateBatchSize:                          tt.fields.CreateBatchSize,
				TranslateError:                           tt.fields.TranslateError,
				ClauseBuilders:                           tt.fields.ClauseBuilders,
				ConnPool:                                 tt.fields.ConnPool,
				Dialector:                                tt.fields.Dialector,
				Plugins:                                  tt.fields.Plugins,
				callbacks:                                tt.fields.callbacks,
				cacheStore:                               tt.fields.cacheStore,
			}
			if err := c.Apply(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_AddError(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if err := db.AddError(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("AddError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Callback(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	tests := []struct {
		name   string
		fields fields
		want   *callbacks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.Callback(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Callback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_DB(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			got, err := db.DB()
			if (err != nil) != tt.wantErr {
				t.Errorf("DB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Debug(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	tests := []struct {
		name   string
		fields fields
		wantTx *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if gotTx := db.Debug(); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Debug() = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}

func TestDB_Get(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			got, got1 := db.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDB_InstanceGet(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			got, got1 := db.InstanceGet(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstanceGet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InstanceGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDB_InstanceSet(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.InstanceSet(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstanceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Session(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		config *Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.Session(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Set(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.Set(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_SetupJoinTable(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		model     interface{}
		field     string
		joinTable interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if err := db.SetupJoinTable(tt.args.model, tt.args.field, tt.args.joinTable); (err != nil) != tt.wantErr {
				t.Errorf("SetupJoinTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Template(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		template string
		param    any
		result   any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//db := &DB{
			//	Config:       tt.fields.Config,
			//	Error:        tt.fields.Error,
			//	RowsAffected: tt.fields.RowsAffected,
			//	Statement:    tt.fields.Statement,
			//	clone:        tt.fields.clone,
			//}
			//if err := db.Template(tt.args.template, tt.args.param, tt.args.result); (err != nil) != tt.wantErr {
			//	t.Errorf("Template() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}

func TestDB_ToSQL(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		queryFn func(tx *DB) *DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.ToSQL(tt.args.queryFn); got != tt.want {
				t.Errorf("ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Use(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		plugin Plugin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if err := db.Use(tt.args.plugin); (err != nil) != tt.wantErr {
				t.Errorf("Use() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_WithContext(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_getInstance(t *testing.T) {
	type fields struct {
		Config       *Config
		Error        error
		RowsAffected int64
		Statement    *Statement
		clone        int
	}
	tests := []struct {
		name   string
		fields fields
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Config:       tt.fields.Config,
				Error:        tt.fields.Error,
				RowsAffected: tt.fields.RowsAffected,
				Statement:    tt.fields.Statement,
				clone:        tt.fields.clone,
			}
			if got := db.getInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpr(t *testing.T) {
	type args struct {
		expr string
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want clause.Expr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Expr(tt.args.expr, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpen(t *testing.T) {
	type args struct {
		dialector Dialector
		opts      []Option
	}
	tests := []struct {
		name    string
		args    args
		wantDb  *DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := Open(tt.args.dialector, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDb, tt.wantDb) {
				t.Errorf("Open() gotDb = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}
