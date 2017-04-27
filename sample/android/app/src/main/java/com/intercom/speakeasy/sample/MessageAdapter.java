package com.intercom.speakeasy.sample;

import android.graphics.Color;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.ViewGroup;

import com.intercom.speakeasy.sample.model.ChannelDataResponse;
import com.intercom.speakeasy.sample.model.Message;
import com.intercom.speakeasy.sample.model.User;

class MessageAdapter extends RecyclerView.Adapter<TextItemViewHolder> {

    private static final int EMPTY_INDICATOR_COUNT = 1;

    private ChannelDataResponse data;

    MessageAdapter(ChannelDataResponse data) {
        this.data = data;
    }

    void setData(ChannelDataResponse data) {
        this.data = data;
    }

    @Override
    public TextItemViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        LayoutInflater inflater = LayoutInflater.from(parent.getContext());
        return new TextItemViewHolder(inflater.inflate(R.layout.text_item, parent, false));
    }

    @Override
    public void onBindViewHolder(final TextItemViewHolder holder, int position) {
        if (data.messages().isEmpty()) {
            bindEmptyView(holder);
        } else {
            bindMessageView(holder, data.messages().get(position));
        }
    }

    private void bindEmptyView(TextItemViewHolder holder) {
        holder.setTextColor(Color.GRAY);
        holder.setText("No messages");
    }

    private void bindMessageView(TextItemViewHolder holder, Message message) {
        holder.setTextColor(Color.BLACK);
        holder.setText(getSenderName(message) + ": " + message.body());
    }

    private String getSenderName(Message message) {
        for (User user : data.users()) {
            if (user.name().equals(message.userName())) {
                return user.name();
            }
        }
        return "";
    }

    @Override
    public int getItemCount() {
        return data.messages().isEmpty()
                ? EMPTY_INDICATOR_COUNT
                : data.messages().size();
    }

}
