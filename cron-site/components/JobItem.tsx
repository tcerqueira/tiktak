import React, { useState } from 'react'
import { CronJob } from '../types/Types';
import { PencilIcon } from '@heroicons/react/solid';
import { TrashIcon } from '@heroicons/react/solid';

interface JobItemProps {
	cronJob: CronJob
}

function JobItem({ cronJob }: JobItemProps) {
	const [editOpen, setEditOpen] = useState(false);

	return (
		<div>
			<div className='flex items-center justify-between'>
				<div>
					<p>Body</p>
					<p>Webhook url</p>
				</div>
				<div className='flex items-center space-x-5'>
					<div className='cronjob__item--div' onClick={() => setEditOpen(!editOpen)}>
						<PencilIcon className='cronjob__item--icon'/>
						<span className='hidden sm:inline'>Edit</span>
					</div>
					<div className='cronjob__item--div'>
						<TrashIcon className='cronjob__item--icon' />
						<span className='hidden sm:inline'>Delete</span>
					</div>
				</div>
			</div>
			{editOpen && 
			<div>
				
			</div>
			}
		</div>
	)
}

export default JobItem