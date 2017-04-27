package com.intercom.speakeasy.sample;

import android.app.Application;

import backend.Backend;
import backend.SampleEngine;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import speakeasy.Server;
import speakeasy.Speakeasy;

public class SpeakeasyApplication extends Application {

    private static SpeakeasyApplication instance;

    public static String USER_NAME = "Sample User";

    private SampleEngine serverEngine;
    private SpeakeasyApi speakeasyApi;

    @Override public void onCreate() {
        super.onCreate();

        instance = this;
        serverEngine = Backend.newEngine();

        final Server speakeasyServer = Speakeasy.newServer(serverEngine);
        Thread serverThread = new Thread(new Runnable() {
            @Override public void run() {
                speakeasyServer.start();
            }
        }, "SpeakeasyServer");
        serverThread.start();

        speakeasyApi = new Retrofit.Builder()
                .addConverterFactory(GsonConverterFactory.create())
                .baseUrl("http://localhost:3000")
                .build()
                .create(SpeakeasyApi.class);
    }

    public static SpeakeasyApplication getInstance() {
        return instance;
    }

    public SpeakeasyApi getSpeakeasyApi() {
        return speakeasyApi;
    }

    public SampleEngine getServerEngine() {
        return serverEngine;
    }

}
