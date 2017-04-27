package com.intercom.speakeasy.sample.model;

import com.google.gson.annotations.SerializedName;

public class NewMessageRequest {

    @SerializedName("user_name") String userName;
    String body;

    public static NewMessageRequest create(String userName, String body) {
        NewMessageRequest request = new NewMessageRequest();
        request.userName = userName;
        request.body = body;
        return request;
    }
}
