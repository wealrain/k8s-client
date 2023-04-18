import * as React from 'react';
import { styled } from '@mui/material/styles';

const Div = styled('div')(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  padding: "15px 10px",
  color: theme.palette.text.primary,
  backgroundColor: "#f5f5f5",
  fontSize: "16px",
  margin: "20px 0px",
}));

export default function DetailTypography(props) {
  return <Div>{props.title}</Div>;
}