import React, { MouseEventHandler, useCallback, useState } from 'react'
import { ClockIcon, PencilIcon } from '@heroicons/react/solid';
import { TrashIcon } from '@heroicons/react/solid';
import TimezoneSelect, { ITimezone } from 'react-timezone-select';
import { CronJob } from '../../types/CronJob';
import { deleteCron, editCron, EditCronPayload } from '../../utils/ApiCalls';
import { useForm } from 'react-hook-form';
import TextArea from '../TextArea';
import TextInput from '../TextInput';
import { cronRegex } from '../../utils/RegEx';

interface EditCronFormData {
	body?: string,
    cron_expression?: string
}

interface JobItemProps {
	cronJob: CronJob,
	onEdit: (id: string, cron: EditCronPayload) => void;
	onDelete: (id: string) => void;
}

function JobItem({ cronJob: { id, webhook_url, webhook_method, body, cron_expression, timezone }, onEdit, onDelete }: JobItemProps) {
	const [editOpen, setEditOpen] = useState(false);
	const [deleteOpen, setDeleteOpen] = useState(false);
	const [selectedTimezone, setSelectedTimezone] = useState<ITimezone>(timezone);
	const { register, handleSubmit, formState: {errors} } = useForm<EditCronFormData>({
		defaultValues: {
			body,
			cron_expression
		}
	});

	const handleDelete: MouseEventHandler<HTMLButtonElement> = evt => {
		onDelete(id);
	};

	const handleEdit = useCallback((data: EditCronFormData) => {
		const timezone = typeof selectedTimezone === 'string' ? selectedTimezone : selectedTimezone.value;
		const payload: EditCronPayload = { ...data, timezone };
		onEdit(id, payload);
	}, [id, selectedTimezone]);

	return (
		<div className='my-2 px-2 rounded-lg border-b-2 border-t-2 border-orange-500 '>
			<div className='flex items-center justify-between'>
				<div className='max-w-[78%]'>
					<p className='break-words'><span className={`font-bold span-${webhook_method}`}>{webhook_method}</span> - {webhook_url}{webhook_method === 'GET' ? `?body=`: ''}</p>
					<HR />
					<span className='px-3 rounded-lg border-l-2 border-r-2 border-orange-500 text-orange-800 font-bold'>{cron_expression}</span>
					<p className='mt-2'>{body}</p>
				</div>
				<div className='flex items-center space-x-5 self-start mt-3'>
					<div className='cronjob__item--div' onClick={() => { setEditOpen(e => !e); setDeleteOpen(false); }}>
						<PencilIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Edit</span>
					</div>
					<div className='cronjob__item--div' onClick={() => { setDeleteOpen(d => !d); setEditOpen(false); }}>
						<TrashIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Delete</span>
					</div>
				</div>
			</div>

			{editOpen && <>
				<HR />
				<form className='flex flex-col p-1' onSubmit={handleSubmit(handleEdit)}>
					<div className='input-container'>
						<TextArea id='body-in' label='Body' {...register('body')}/>
					</div>
					<div className={`input-container ${errors.cron_expression && 'input-error'}`}>
						<TextInput id='schedule-in' label='Schedule' placeholder='* * * * *' {...register('cron_expression', { required: true, pattern: cronRegex })}/>
					</div>
					{ errors.cron_expression && <p className='error-message'>Invalid CRON expression.</p> }
					<div className='input-container'>
						<label htmlFor='timezone-in'>Timezone</label>
						<TimezoneSelect value={selectedTimezone} onChange={setSelectedTimezone} />
					</div>
					<button type='submit' className='submit-btn bg-blue-200 mt-4'>
						<ClockIcon className='h-10 w-10 text-orange-600' />
					</button>
				</form>
				</>
			}
			{deleteOpen && <>
				<HR />
				<div className='flex justify-center items-center space-x-5 py-3'>
					<button className='delete-prompt-btns bg-red-300' onClick={handleDelete}>Delete</button>
					<button className='delete-prompt-btns bg-gray-300' onClick={() => setDeleteOpen(d => !d)}>Cancel</button>
				</div>
				</>
			}
		</div>
	)
}

export default JobItem

interface HRProps {
	className?: string;
}

const HR = ({ className }: HRProps) => {
	return (
		<hr className={'border-orange-300'}/>
	);
}