package com.intercom.speakeasy.sample.model;

import com.google.gson.annotations.SerializedName;

public class NewChannelRequest {

    @SerializedName("user_name") String userName;
    @SerializedName("channel_name") String channelName;

    public static NewChannelRequest create(String userName, String channelName) {
        NewChannelRequest request = new NewChannelRequest();
        request.userName = userName;
        request.channelName = channelName;
        return request;
    }
}
