package com.intercom.speakeasy.sample;

import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.RecyclerView;
import android.support.v7.widget.Toolbar;
import android.view.MenuItem;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import com.intercom.speakeasy.sample.model.ChannelDataResponse;
import com.intercom.speakeasy.sample.model.Message;
import com.intercom.speakeasy.sample.model.NewMessageRequest;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class MessagesActivity extends AppCompatActivity {

    static final String KEY_CHANNEL_NAME = "channel_name";

    private SpeakeasyApi speakeasyApi;
    private MessageAdapter adapter;
    private String channelName;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_messages);

        channelName = getIntent().getStringExtra(KEY_CHANNEL_NAME);

        Toolbar toolbar = (Toolbar) findViewById(R.id.detail_toolbar);
        setSupportActionBar(toolbar);

        ActionBar actionBar = getSupportActionBar();
        actionBar.setDisplayHomeAsUpEnabled(true);
        actionBar.setTitle(channelName);

        adapter = new MessageAdapter(ChannelDataResponse.emptyResponse());
        RecyclerView recyclerView = (RecyclerView) findViewById(R.id.message_list);
        recyclerView.setAdapter(adapter);

        speakeasyApi = SpeakeasyApplication.getInstance().getSpeakeasyApi();
        fetchMessages();

        findViewById(R.id.send_button).setOnClickListener(new View.OnClickListener() {
            @Override public void onClick(View v) {
                sendMessage();
            }
        });
    }

    private void fetchMessages() {
        speakeasyApi.fetchChannelData(channelName).enqueue(new Callback<ChannelDataResponse>() {
            @Override public void onResponse(Call<ChannelDataResponse> call, Response<ChannelDataResponse> response) {
                if (response.isSuccessful()) {
                    adapter.setData(response.body());
                    adapter.notifyDataSetChanged();
                } else {
                    showError("Failed to fetch channel data");
                }
            }

            @Override public void onFailure(Call<ChannelDataResponse> call, Throwable t) {
                showError("Failed to fetch channel data");
            }

        });
    }

    private void showError(String text) {
        Toast.makeText(this, text, Toast.LENGTH_SHORT).show();
    }

    private void sendMessage() {
        final EditText sendTextView = (EditText) findViewById(R.id.message_edit_text);

        String messageBody = sendTextView.getText().toString();
        NewMessageRequest request = NewMessageRequest.create(SpeakeasyApplication.USER_NAME, messageBody);
        speakeasyApi.sendMessage(channelName, request).enqueue(new Callback<Message>() {
            @Override public void onResponse(Call<Message> call, Response<Message> response) {
                if (response.isSuccessful()) {
                    fetchMessages();
                } else {
                    showError("Failed to send message");
                }
                sendTextView.setText("");
            }

            @Override public void onFailure(Call<Message> call, Throwable t) {
                showError("Failed to send message");
            }

        });
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                navigateUpTo(new Intent(this, ChannelsActivity.class));
                return true;
            default:
                return super.onOptionsItemSelected(item);
        }
    }
}
