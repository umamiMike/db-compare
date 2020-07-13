import React from "react";

interface Props {
  name: string;
  value?: string;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

const TextArea = (props: Props) => {
  /* TODO:
   */
  const initProps = {
    className: "form-component w-full border border-grey",
    style: { height: "15em" },
  };
  const elProps = { ...initProps, ...props };

  return <textarea {...elProps} />;
};

export default TextArea;
