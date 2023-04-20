import * as React from 'react';
import { styled } from '@mui/material/styles';

const Div = styled('span')(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  padding: "10px",
  color: theme.palette.text.primary,
  fontSize: "14px",
  display: "inline-block",
}));

export default function DetailTypography(props) {
  return <Div>{props.children}</Div>;
}
