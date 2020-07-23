import React from "react";

interface Props {
  name: string;
  value?: string;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export default function TextArea(props: Props) {
  const initProps = {
    style: { height: "15em" },
  };
  const elProps = { ...initProps, ...props };

  return <textarea {...elProps} />;
}

