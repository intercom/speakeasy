package com.intercom.speakeasy.sample;

import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import java.util.List;

class ChannelAdapter extends RecyclerView.Adapter<TextItemViewHolder> {

    private final ChannelClickListener listener;

    private List<String> channels;

    ChannelAdapter(List<String> channels, ChannelClickListener listener) {
        this.channels = channels;
        this.listener = listener;
    }

    void setData(List<String> channels) {
        this.channels = channels;
    }

    @Override
    public TextItemViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        LayoutInflater inflater = LayoutInflater.from(parent.getContext());
        return new TextItemViewHolder(inflater.inflate(R.layout.text_item, parent, false));
    }

    @Override
    public void onBindViewHolder(final TextItemViewHolder holder, int position) {
        final String channel = channels.get(position);
        holder.setText(channel);
        holder.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                listener.onChannelClicked(channel);
            }
        });
    }

    @Override
    public int getItemCount() {
        return channels.size();
    }

}
