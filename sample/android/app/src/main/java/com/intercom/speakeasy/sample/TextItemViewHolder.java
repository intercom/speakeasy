package com.intercom.speakeasy.sample;

import android.support.annotation.ColorInt;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.TextView;

class TextItemViewHolder extends RecyclerView.ViewHolder {
    private final TextView textView;

    TextItemViewHolder(View view) {
        super(view);
        textView = (TextView) view.findViewById(R.id.text_view);
    }

    void setTextColor(@ColorInt int color) {
        textView.setTextColor(color);
    }

    void setText(CharSequence text) {
        textView.setText(text);
    }

    void setOnClickListener(View.OnClickListener onClickListener) {
        itemView.setOnClickListener(onClickListener);
    }
}
