import React from "react";
import CodeMirror from "@uiw/react-codemirror";
import { githubLight } from "@uiw/codemirror-theme-github";

export default function ConfigReader(props) {
  const {code} = props;
  console.log(code);
  return (
    <CodeMirror
      width="560px"
      height="400px"
      value={code}
      theme={githubLight}
      readOnly={true}
    />
  );
}