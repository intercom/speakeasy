package com.intercom.speakeasy.sample.model;

import com.google.gson.annotations.SerializedName;

public class Message {

    String id;
    @SerializedName("user_name") String userName;
    @SerializedName("created_at") long createdAt;
    String body;

    public String id() {
        return id;
    }

    public String userName() {
        return userName;
    }

    public long createdAt() {
        return createdAt;
    }

    public String body() {
        return body;
    }
}
