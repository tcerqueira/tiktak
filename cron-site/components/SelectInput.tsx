import React from 'react'

interface SelectInputProps {
    id: string;
    name: string;
    label: string;
    onChange?: ((opt: React.ChangeEvent<HTMLSelectElement>) => void)
    children?: JSX.Element | JSX.Element[]
}

const SelectInput = React.forwardRef<HTMLSelectElement, SelectInputProps>(({ id, name, label, onChange, children, ...props }, ref) => (
    <>
        <label htmlFor={id}>{label}</label>
        <select ref={ref} id={id} onChange={onChange} {...props}>
            {children}
        </select>
    </>
))


export default SelectInput