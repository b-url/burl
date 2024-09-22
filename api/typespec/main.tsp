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
    @segment("users")
    @visibility("read")
    @doc("Resource name (e.g., users/{user_id})")
    name: string; 
    
    @key("userId")
    @visibility("read")
    uid: string; 

    @visibility("read")
    createTime: utcDateTime; 

    @visibility("read")
    updateTime: utcDateTime;
}


@resource("collections")
@parentResource(User)
model Collection {
    @segment("collections")
    @visibility("read")
    @doc("Resource name (e.g., users/{user_id}/collections/{collection_id})")
    name: string; 
    
    @key("collectionId")
    @docFromComment
    @visibility("read")
    @doc("System-generated unique identifier for the collection.")
    uid: string; 
    
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
    @visibility("read")
    @doc("Resource name (e.g., users/{user_id}/collections/{collection_id}/bookmarks/{bookmark_id})")
    name: string; 
    
    @key("bookmarkId")
    @docFromComment
    @visibility("read")
    @doc("System-generated unique identifier for the bookmark.")
    uid: string; 

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

@autoRoute
interface Collections {
    @doc("Create a collection")
    op create is CollectionOperations.create;

    @doc("Get a collection")
    op read is CollectionOperations.get;

    @doc("Update a collection")
    op update is CollectionOperations.update;
}

