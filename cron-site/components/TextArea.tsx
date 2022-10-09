import React from 'react'

interface TextAreaProps {
    id: string;
    label: string;
    cols?: number,
    rows?: number;
    defaultValue?: string;
    [props: string]: unknown;
}

const TextArea = React.forwardRef<HTMLTextAreaElement, TextAreaProps>(({ id, label, cols, rows, defaultValue, ...props }, ref) => (
    <>
        <label htmlFor={id}>{label}</label>
        <textarea ref={ref} id={id} cols={cols || 30} rows={rows || 4} defaultValue={defaultValue || ''} {...props} />
    </>
))

export default TextArea