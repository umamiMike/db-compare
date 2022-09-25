import React, { useState } from 'react';

interface Props {
  creds: DatasourceCredentials;
  onSubmit: any;
}

export default function DatasourceList(props: Props) {
  const { creds, onSubmit } = props;
  const submitFormData = onSubmit;
  const [formdata, setFormData] = useState(creds);

  return (
    <div className='ds-list'>
      <form>
        <button
          onClick={e => {
            e.preventDefault();
            submitFormData(formdata);
          }}
        >
          {' '}
          submit
        </button>
      </form>
    </div>
  );
}
