import React from 'react'

interface SelectInputProps {
    label: string;
    onChange?: ((opt: React.ChangeEvent<HTMLSelectElement>) => void);
    children?: JSX.Element | JSX.Element[];
    [props: string]: unknown;
}

const SelectInput = React.forwardRef<HTMLSelectElement, SelectInputProps>(({ label, onChange, children, ...props }, ref) => (
    <>
        <label>{label}</label>
        <select ref={ref} onChange={onChange} {...props}>
            {children}
        </select>
    </>
))


export default SelectInput