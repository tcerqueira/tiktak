import React from 'react'

interface InputProps {
  id: string;
  label: string;
  placeholder?: string;
  [props: string]: unknown;
}

const TextInput = React.forwardRef<HTMLInputElement, InputProps>(({ id, label, placeholder, ...props }, ref) => (
  <>
    <label htmlFor={id}>{label}</label>
    <input ref={ref} id={id} type='text' placeholder={placeholder || ''} {...props} />
  </>
));


export default TextInput