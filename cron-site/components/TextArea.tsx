import React from 'react'

interface TextAreaProps {
    id: string;
    name: string;
    label: string;
    cols?: number,
    rows?: number;
    defaultValue?: string;
}

const TextArea = React.forwardRef<HTMLTextAreaElement, TextAreaProps>(({ id, name, label, cols, rows, defaultValue, ...props }, ref) => (
    <>
        <label htmlFor={id}>{label}</label>
        <textarea ref={ref} id={id} cols={cols || 30} rows={rows || 4} defaultValue={defaultValue || ''} {...props} />
    </>
))

export default TextArea