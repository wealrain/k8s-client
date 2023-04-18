import React from "react";
import { linter, lintGutter } from "@codemirror/lint";
import CodeMirror from "@uiw/react-codemirror";
import * as yamlMode from "@codemirror/legacy-modes/mode/yaml";
import { StreamLanguage } from "@codemirror/language";
import parser from "js-yaml";
import { githubLight } from "@uiw/codemirror-theme-github";

const yaml = StreamLanguage.define(yamlMode.yaml);

const yamlLinter = linter((view) => {
  const diagnostics = [];

  try {
    parser.load(view.state.doc);
  } catch (e) {
    const loc = e.mark;
    const from = loc ? loc.position : 0;
    const to = from;
    const severity = "error";

    diagnostics.push({
      from,
      to,
      message: e.message,
      severity
    });
  }

  return diagnostics;
});

export default function YamlEditor(props) {

  return (
    <CodeMirror
      width="100%"
      value={props.value}
      theme={githubLight}
      extensions={[yaml, lintGutter(), yamlLinter]}
      onChange={(editor, data, value) => {
        
        props.onChange && props.onChange(editor);
      }}
    />
  );
}