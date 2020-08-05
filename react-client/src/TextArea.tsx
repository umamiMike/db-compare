import React from "react";

interface Props {
  name: string;
  value?: string;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export default function TextArea(props: Props) {
  const initProps = {};
  const elProps = { ...initProps, ...props };

  return <textarea {...elProps} />;
}

