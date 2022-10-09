import React, { MouseEventHandler, useCallback, useState } from 'react'
import { ClockIcon, PencilIcon } from '@heroicons/react/solid';
import { TrashIcon } from '@heroicons/react/solid';
import TimezoneSelect, { ITimezone } from 'react-timezone-select';
import { CronJob } from '../../types/CronJob';
import { deleteCron } from '../../utils/ApiCalls';

interface JobItemProps {
	cronJob: CronJob
}

function JobItem({ cronJob: { id, webhook_url, webhook_method, body, cron_expression, timezone } }: JobItemProps) {
	const [editOpen, setEditOpen] = useState(false);
	const [deleteOpen, setDeleteOpen] = useState(false);
	const [selectedTimezone, setSelectedTimezone] = useState<ITimezone>(timezone);

	const onDelete: MouseEventHandler<HTMLButtonElement> = useCallback(async evt => {
		try {
			await deleteCron(id);
		} catch (err) {
			console.error(err);
		}
	}, []);

	return (
		<div className='my-2 px-2 rounded-lg border-b-2 border-t-2 border-orange-500 '>
			<div className='flex items-center justify-between'>
				<div className='max-w-[78%]'>
					<p><span className={`span-${webhook_method}`}>{webhook_method}</span> - {webhook_url}{webhook_method === 'GET' ? `?body=`: ''}</p>
					<HR />
					<p className='mt-2'>{body}</p>
				</div>
				<div className='flex items-center space-x-5 self-start mt-2'>
					<div className='cronjob__item--div' onClick={() => setEditOpen(e => !e)}>
						<PencilIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Edit</span>
					</div>
					<div className='cronjob__item--div' onClick={() => setDeleteOpen(d => !d)}>
						<TrashIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Delete</span>
					</div>
				</div>
			</div>

			{editOpen && <>
				<HR />
				<form className='p-1'>
					<div className='input-container'>
						<label htmlFor='body-in'>Body</label>
						<textarea id="body-in" cols={30} rows={4} defaultValue={body} />
					</div>
					<div className='input-container'>
						<label htmlFor='schedule-in'>Schedule</label>
						<input id='schedule-in' defaultValue={cron_expression} type='text' placeholder='* * * * *' />
					</div>
					<div className='input-container'>
						<label htmlFor='timezone-in'>Timezone</label>
						<TimezoneSelect value={selectedTimezone} onChange={setSelectedTimezone} />
					</div>
					<button type='submit' className='flex justify-center items-center md:w-[60%] w-[100%] rounded-lg mx-auto cursor-pointer hover:backdrop-brightness-90'>
						<ClockIcon className='h-10 w-10 text-orange-600' />
					</button>
				</form>
				</>
			}
			{deleteOpen && <>
				<HR />
				<div className='flex justify-center items-center space-x-5 py-3'>
					<button className='delete-prompt-btns bg-red-300' onClick={onDelete}>Delete</button>
					<button className='delete-prompt-btns bg-gray-300' onClick={() => setDeleteOpen(d => !d)}>Cancel</button>
				</div>
				</>
			}
		</div>
	)
}

export default JobItem

const HR = () => {
	return (
		<hr className='border-orange-300'/>
	);
}