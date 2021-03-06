package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AccountPlatform holds the schema definition for the AccountPlatform entity.
type AccountPlatform struct {
	ent.Schema
}

// Fields of the AccountPlatform.
func (AccountPlatform) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("uid").Default(0).Positive().Comment("账号 id"),
		field.String("platform_id").MaxLen(30).Default("").Comment("平台 id"),
		field.String("platform_token").MaxLen(32).Default("").Comment("平台 access_token"),
		field.String("nickname").MaxLen(60).Default("").Comment(""),
		field.String("avatar").MaxLen(255).Default("").Comment(""),
		field.Int8("type").Default(0).Comment("平台类型 0: 未知，1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter"),
	}
}

// Edges of the AccountPlatform.
func (AccountPlatform) Edges() []ent.Edge {
	return nil
}
func (AccountPlatform) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uid", "platform_id"),
	}
}
func (AccountPlatform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
