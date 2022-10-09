import React from 'react'
import { CronJob } from '../../types/CronJob';
import JobItem from './JobItem'

interface JobListProps {
  cronList: CronJob[];
  isLoading: boolean;
  error?: string;
}

function JobList({ cronList, isLoading, error }: JobListProps) {
  return (
    <div>
      <h1 className='text-center'>{error ? error : isLoading ? "Loading..." : "CRONS"}</h1>
      {!cronList.length && <p>No CRON jobs...</p>}
      <ul>
        {!isLoading && cronList?.map((item, i) => (
          <li key={item.id}>
            <JobItem cronJob={item} />
          </li>
        ))
      }
      </ul>
    </div>
  )
}

export default JobList