import { Typography } from '@mui/material';
import * as React from 'react';

const style = (content) => {
    let color = "#3c3c3c";
    // 判断是否包含关键字
    if (content.indexOf("running") !== -1) {
        color = "#00bfa5";
    }
    if (content.indexOf("pending") !== -1) {
        color = "#ff9800";
    }
    if (content.indexOf("failed") !== -1) {
        color = "#f44336";
    }
    if (content.indexOf("succeeded") !== -1) {
        color = "#00bfa5";
    }
    if (content.indexOf("unknown") !== -1) {
        color = "#3c3c3c";
    }

    return {
        color: color,
        fontSize: "14px",
        padding: "10px",
        display: "inline-block",
    }
}

export default function StatusTypography(props) {
  return (
    <Typography style={style(props.children.toLowerCase())}>
        {props.children}
    </Typography>)
}