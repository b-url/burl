import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";
import "@typespec/versioning";

import "./error.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;


@service({
    title: "burl API",
})
@TypeSpec.Versioning.versioned(Versions)
namespace BurlService;

enum Versions {
    v1,
}

@resource("users")
model User {
    @doc("System-generated unique identifier for the user.")
    @encode("uuid", string)
    @key("userId")
    @segment("users")
    @TypeSpec.OpenAPI.extension("x-go-type", "uuid.UUID")
    @visibility("read")
    id: string; 

    @visibility("read")
    createTime: utcDateTime;
    @visibility("read")
    updateTime: utcDateTime;
}


@resource("collections")
@parentResource(User)
model Collection {
    @doc("System-generated unique identifier for the collection.")
    @encode("uuid", string)
    @key("collectionId")
    @segment("collections")
    @TypeSpec.OpenAPI.extension("x-go-type", "uuid.UUID")
    @visibility("read")
    id: string;

    @doc("System-generated unique identifier for the parent collection.")
    @encode("uuid", string)
    @TypeSpec.OpenAPI.extension("x-go-type", "uuid.UUID")
    @visibility("read", "create", "update")
    parentId?: string;
    
    @visibility("read", "create", "update")
    name: string;

    @visibility("read")
    createTime: utcDateTime; 
    @visibility("read")
    updateTime: utcDateTime;
}



@doc("Bookmark is a resource that represents a saved URL.")
@resource("bookmarks")
@parentResource(Collection)
model Bookmark {
    @doc("System-generated unique identifier for the bookmark.")
    @encode("uuid", string)
    @key("bookmarkId")
    @segment("bookmarks")
    @TypeSpec.OpenAPI.extension("x-go-type", "uuid.UUID")
    id: string;
    
    @doc("System-generated unique identifier for the parent collection.")
    @encode("uuid", string)
    @TypeSpec.OpenAPI.extension("x-go-type", "uuid.UUID")
    @visibility("read")
    parentId?: string;

    @visibility("read", "create")
    url: string; 

    @visibility("read", "create", "update")
    title: string; 

    @visibility("read", "create", "update")
    tags: string[];

    @visibility("read")
    createTime: utcDateTime; 
    @visibility("read")
    updateTime: utcDateTime;
}

alias BookmarkOperations = TypeSpec.Rest.Resource.ResourceOperations<Bookmark, Error>;

@autoRoute
interface Bookmarks {
    @doc("Create a bookmark")
    op create is BookmarkOperations.create;

    @doc("Get a bookmark")
    op read is BookmarkOperations.get;

    @doc("Update a bookmark")
    op update is BookmarkOperations.update;
}

alias CollectionOperations = TypeSpec.Rest.Resource.ResourceOperations<Collection, Error>;
alias CollectionParams = TypeSpec.Rest.Resource.ResourceParameters<Collection>;

model CollectionChildren {
    bookmarks: Bookmark[];
    collections: Collection[];
}

@autoRoute
interface Collections {
    @doc("Create a collection")
    op create is CollectionOperations.create;

    @doc("Get a collection")
    op read is CollectionOperations.get;

    @doc("Update a collection")
    op update is CollectionOperations.update;

    @segment("children")
    @doc("List all the children of a collection, including bookmarks and sub-collections.")
    op listChildren(...CollectionParams): CollectionChildren[] | Error;
}

