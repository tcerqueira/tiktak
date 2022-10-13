import React, { useCallback } from 'react'
import { CronJob } from '../../types/CronJob';
import { deleteCron, editCron, EditCronPayload } from '../../utils/ApiCalls';
import JobItem from './JobItem'

interface JobListProps {
	cronList: CronJob[];
	isLoading: boolean;
	error?: string;
	fetchList: () => void;
}

function JobList({ cronList, isLoading, error, fetchList }: JobListProps) {
	const onDelete = useCallback(async (id: string) => {
		try {
			await deleteCron(id);
			fetchList();
		} catch (err) {
			console.error(err);
		}
	}, []);

	const onEdit = useCallback(async (id: string, cron: EditCronPayload) => {
		try {
			await editCron(id, cron);
			fetchList();
		} catch (err) {
			console.error(err);
		}
	}, []);

	return (
		<div>
			<h1 className='text-center font-bold'>{error ? error : isLoading ? "Loading..." : "CRONS"}</h1>
			{!cronList.length && <p>No CRON jobs...</p>}
			<ul>
				{!isLoading && cronList?.map((item) => (
					<li key={item.id}>
						<JobItem cronJob={item} onEdit={onEdit} onDelete={onDelete}/>
					</li>
				))
				}
			</ul>
		</div>
	)
}

export default JobList