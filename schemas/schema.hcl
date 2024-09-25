schema "public" {
  comment = "the burl db"
}

table "users" {
  schema = schema.public

  column "id" {
    type = int
    identity {
      generated = ALWAYS
    }
  }

  column "username" {
    type = text
    null = true
  }

  column "email" {
    type = text
    null = true
  }

  column "created_at" {
    type = timestamp
  }

  column "updated_at" {
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_username" {
    columns = [
      column.username
    ]

  }

  index "idx_email" {
    columns = [
      column.email
    ]
  }
}

table "collections" {
  schema = schema.public

  column "id" {
    type = int
    identity {
      generated = ALWAYS
    }
  }

  column "parent_id" {
    type = bigint
    null = true
  }

  foreign_key "parent_id_fk" {
    columns     = [column.parent_id]
    ref_columns = [table.collections.column.id]
    on_delete   = CASCADE
  }

  column "user_id" {
    type = bigint
  }

  foreign_key "user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete   = CASCADE
  }

  column "name" {
    type = text
    null = true
  }

  column "created_at" {
    type = timestamp
  }

  column "updated_at" {
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }
}

table "bookmarks" {
  schema = schema.public

  column "id" {
    type = int
    identity {
      generated = ALWAYS
    }
  }

  column "collection_id" {
    type = bigint
  }

  foreign_key "collection_id_fk" {
    columns     = [column.collection_id]
    ref_columns = [table.collections.column.id]
    on_delete   = CASCADE
  }

  column "user_id" {
    type = bigint
  }

  foreign_key "user_id_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_delete   = CASCADE
  }

  column "url" {
    type = text
    null = true
  }

  column "title" {
    type = text
    null = true
  }

  column "description" {
    type = text
  }

  column "created_at" {
    type = timestamp
  }

  column "updated_at" {
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }
}

table "tags" {
  schema = schema.public

  column "id" {
    type = int
    identity {
      generated = ALWAYS
    }
  }

  column "name" {
    type = text
    null = true
  }

  primary_key {
    columns = [column.id]
  }
}

table "bookmark_tags" {
  schema = schema.public

  column "bookmark_id" {
    type = int
  }

  foreign_key "bookmark_fk" {
    columns     = [column.bookmark_id]
    ref_columns = [table.bookmarks.column.id]
    on_delete   = CASCADE
  }

  column "tag_id" {
    type = int
  }

  foreign_key "tag_fk" {
    columns     = [column.tag_id]
    ref_columns = [table.tags.column.id]
    on_delete   = CASCADE
  }

  primary_key {
    columns = [column.bookmark_id, column.tag_id]
  }
}