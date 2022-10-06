import React from 'react'

interface InputProps {
  id: string;
  name: string;
  label: string;
  placeholder?: string;
}

const TextInput = React.forwardRef<HTMLInputElement, InputProps>(({ id, name, label, placeholder, ...props }, ref) => (
  <>
    <label htmlFor={id}>Schedule</label>
    <input ref={ref} id={id} type='text' name={name} placeholder={placeholder || ''} {...props} />
  </>
));


export default TextInput