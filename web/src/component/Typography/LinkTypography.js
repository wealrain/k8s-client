import * as React from 'react';
import { styled } from '@mui/material/styles';

const Div = styled('span')(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  padding: "10px",
  color: " #0079F4",
  fontSize: "14px",
  display: "inline-block",
}));

export default function LinkTypography(props) {
  return <Div>{props.children}</Div>;
}