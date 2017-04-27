package com.intercom.speakeasy.sample;

import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AlertDialog;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.RecyclerView;
import android.support.v7.widget.Toolbar;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import com.intercom.speakeasy.sample.model.ChannelListResponse;
import com.intercom.speakeasy.sample.model.NewChannelRequest;

import java.util.ArrayList;

import backend.SampleEngine;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class ChannelsActivity extends AppCompatActivity {

    private ChannelAdapter adapter;
    private SampleEngine serverEngine;
    private SpeakeasyApi speakeasyApi;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_channels);

        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);
        toolbar.setTitle(getTitle());

        adapter = new ChannelAdapter(new ArrayList<String>(), new ChannelClickListener() {
            @Override public void onChannelClicked(String channel) {
                openChannel(channel);
            }
        });
        RecyclerView recyclerView = (RecyclerView) findViewById(R.id.channel_list);
        recyclerView.setAdapter(adapter);

        speakeasyApi = SpeakeasyApplication.getInstance().getSpeakeasyApi();
        serverEngine = SpeakeasyApplication.getInstance().getServerEngine();
        fetchChannels();

        findViewById(R.id.create_channel_fab).setOnClickListener(new View.OnClickListener() {
            @Override public void onClick(View v) {
                Context context = v.getContext();
                final EditText input = new EditText(context);
                new AlertDialog.Builder(context)
                        .setTitle("Create Channel")
                        .setView(input)
                        .setNegativeButton("Cancel", null)
                        .setPositiveButton("Create", new DialogInterface.OnClickListener() {
                            @Override public void onClick(DialogInterface dialog, int which) {
                                createChannel(input.getText().toString());
                            }
                        })
                        .show();
            }
        });
    }

    private void openChannel(String channel) {
        startActivity(new Intent(this, MessagesActivity.class)
                .putExtra(MessagesActivity.KEY_CHANNEL_NAME, channel));
    }

    private void fetchChannels() {
        speakeasyApi.fetchAllChannels().enqueue(new Callback<ChannelListResponse>() {
            @Override public void onResponse(Call<ChannelListResponse> call, Response<ChannelListResponse> response) {
                if (response.isSuccessful()) {
                    adapter.setData(response.body().channels());
                    adapter.notifyDataSetChanged();
                } else {
                    showError("Failed to create channel");
                }
            }

            @Override public void onFailure(Call<ChannelListResponse> call, Throwable t) {
                showError("Failed to fetch channels");
            }
        });
    }

    private void showError(String text) {
        Toast.makeText(this, text, Toast.LENGTH_SHORT).show();
    }

    private void createChannel(String name) {
        speakeasyApi.createChannel(NewChannelRequest.create(SpeakeasyApplication.USER_NAME, name))
                .enqueue(new Callback<Void>() {
                    @Override public void onResponse(Call<Void> call, Response<Void> response) {
                        if (response.isSuccessful()) {
                            fetchChannels();
                        } else {
                            showError("Failed to create channel");
                        }
                    }

                    @Override public void onFailure(Call<Void> call, Throwable t) {
                        showError("Failed to create channel");
                    }
                });
    }

    @Override public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.main, menu);
        return true;
    }

    @Override public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case R.id.action_reset:
                serverEngine.reset();
                showError("Server has been reset");
                break;
            case R.id.action_enable_failure:
                serverEngine.setAlwaysFail(true);
                showError("Server failure has been enabled");
                break;
            case R.id.action_disable_failure:
                serverEngine.setAlwaysFail(false);
                showError("Server failure has been disabled");
                break;
            default:
        }
        return super.onOptionsItemSelected(item);
    }

}
