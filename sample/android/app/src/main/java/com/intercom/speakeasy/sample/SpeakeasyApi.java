package com.intercom.speakeasy.sample;

import com.intercom.speakeasy.sample.model.ChannelDataResponse;
import com.intercom.speakeasy.sample.model.ChannelListResponse;
import com.intercom.speakeasy.sample.model.Message;
import com.intercom.speakeasy.sample.model.NewChannelRequest;
import com.intercom.speakeasy.sample.model.NewMessageRequest;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.Path;

interface SpeakeasyApi {

    @GET("channels/{channelName}") Call<ChannelDataResponse> fetchChannelData(@Path("channelName") String channelName);

    @POST("channels/{channelName}/reply") Call<Message> sendMessage(
            @Path("channelName") String channelName,
            @Body NewMessageRequest request
    );

    @GET("channels") Call<ChannelListResponse> fetchAllChannels();

    @POST("channels") Call<Void> createChannel(
            @Body NewChannelRequest request
    );

}
