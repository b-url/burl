import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
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
    @key("userId")
    @segment("users")
    @visibility("read")
    @doc("System-generated unique identifier for the user.")
    id: string; 

    @visibility("read")
    createTime: utcDateTime; 

    @visibility("read")
    updateTime: utcDateTime;
}


@resource("collections")
@parentResource(User)
model Collection {
    @segment("collections")
    @key("collectionId")
    @visibility("read")
    @doc("System-generated unique identifier for the collection.")
    id: string;

    @visibility("read", "create", "update")
    @doc("System-generated unique identifier for the parent collection.")
    parentId?: string;
    
    @visibility("read")
    parentCollection?: string;

    @visibility("read", "create", "update")
    displayName: string;
    
    @visibility("read", "create", "update")
    description?: string;

    @visibility("read")
    createTime: utcDateTime; 

    @visibility("read")
    updateTime: utcDateTime;
}

@doc("Bookmark is a resource that represents a saved URL.")
@resource("bookmarks")
@parentResource(Collection)
model Bookmark {
    @segment("bookmarks")
    @key("bookmarkId")
    @visibility("read")
    @doc("System-generated unique identifier for the bookmark.")
    id: string;
    
    @visibility("read", "create", "update")
    @doc("System-generated unique identifier for the parent collection.")
    parentId?: string;

    @visibility("read", "create")
    url: string; 

    @visibility("read", "create", "update")
    displayName?: string; 

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

