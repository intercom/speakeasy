package com.intercom.speakeasy.sample.model;

import com.google.gson.annotations.SerializedName;

public class User {

    String name;
    @SerializedName("image_url") String imageUrl;

    public String name() {
        return name;
    }

    public String imageUrl() {
        return imageUrl;
    }
}
