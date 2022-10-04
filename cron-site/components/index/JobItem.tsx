import React, { useState } from 'react'
import { ClockIcon, PencilIcon } from '@heroicons/react/solid';
import { TrashIcon } from '@heroicons/react/solid';
import TimezoneSelect, { ITimezone } from 'react-timezone-select';
import { CronJob } from '../../types/CronJob';

interface JobItemProps {
	cronJob: CronJob
}

function JobItem({ cronJob }: JobItemProps) {
	const [editOpen, setEditOpen] = useState(false);
	const [selectedTimezone, setSelectedTimezone] = useState<ITimezone>(cronJob.timezone);

	return (
		<div className='my-2 px-2 rounded-lg border-b-2 border-t-2 border-orange-500 '>
			<div className='flex items-center justify-between'>
				<div>
					<p>{cronJob.body}</p>
					<p>{cronJob.webhook_url}</p>
				</div>
				<div className='flex items-center space-x-5'>
					<div className='cronjob__item--div' onClick={() => setEditOpen(!editOpen)}>
						<PencilIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Edit</span>
					</div>
					<div className='cronjob__item--div'>
						<TrashIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Delete</span>
					</div>
				</div>
			</div>

			{editOpen &&
				<form className='p-1'>
					<div className='input-container'>
						<label htmlFor='body-in'>Body</label>
						<textarea id="body-in" cols={30} rows={4} defaultValue={cronJob.body} />
					</div>
					<div className='input-container'>
						<label htmlFor='schedule-in'>Schedule</label>
						<input id='schedule-in' defaultValue={cronJob.expression} type='text' placeholder='* * * * *' />
					</div>
					<div className='input-container'>
						<label htmlFor='timezone-in'>Timezone</label>
						<TimezoneSelect value={selectedTimezone} onChange={setSelectedTimezone} />
					</div>
					<div className='flex justify-center items-center md:w-[60%] w-[100%] rounded-lg mx-auto cursor-pointer hover:backdrop-brightness-90'
						onClick={() => console.log('edit')}>
						<ClockIcon className='h-10 w-10 text-orange-600' />
						<button type='submit' hidden />
					</div>
				</form>
			}
		</div>
	)
}

export default JobItem