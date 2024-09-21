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
    name: string; // Resource name (e.g., users/{user_id})
    
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
    @key("collectionId")
    @segment("collections")
    @visibility("read")
    name: string; // Resource name (e.g., users/{user_id}/collections/{collection_id})
    
    @visibility("read")
    uid: string; 

    @visibility("read", "create", "update")
    displayName: string; 

    @visibility("read")
    createTime: utcDateTime; 

    @visibility("read")
    updateTime: utcDateTime;
}

@doc("Bookmark is a resource that represents a saved URL.")
@resource("bookmarks")
@parentResource(Collection)
model Bookmark {
    @key
    @segment("bookmarks")
    @visibility("read")
    name: string; // Resource name (e.g., users/{user_id}/collections/{collection_id}/bookmarks/{bookmark_id})
    
    @visibility("read")
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


alias CreateBookmark = TypeSpec.Rest.Resource.ResourceCreate<Bookmark, Error>;

@autoRoute
interface Bookmarks {
    @doc("Create a bookmark")
    op create is CreateBookmark.create;
}