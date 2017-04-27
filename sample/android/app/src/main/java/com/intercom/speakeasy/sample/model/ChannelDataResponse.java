package com.intercom.speakeasy.sample.model;

import java.util.Collections;
import java.util.List;

public class ChannelDataResponse {

    List<User> users;
    List<Message> messages;

    public List<Message> messages() {
        return messages;
    }

    public List<User> users() {
        return users;
    }

    public static ChannelDataResponse emptyResponse() {
        ChannelDataResponse response = new ChannelDataResponse();
        response.users = Collections.emptyList();
        response.messages = Collections.emptyList();
        return response;
    }
}
