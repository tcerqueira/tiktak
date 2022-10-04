import React, { useEffect, useRef, useState, memo } from 'react'
import useCronList from '../../hooks/useCronList';
import { CronJob } from '../../types/CronJob'
import JobItem from './JobItem'
import { v4 as uuid } from "uuid"

const mock: CronJob = {
  id: "578427823",
  webhook_url: 'https://localhost:3000/api/',
  webhook_method: 'POST',
  body: 'ITS TIME MF',
  expression: '* * * * *',
  timezone: 'Europe/Lisbon'
};

const mockList: CronJob[] = [
  { ...mock }, { ...mock }, { ...mock }
]

function JobList() {
  const { cronList, isLoading, error } = useCronList();

  return (
    <div>
      <h1 className='text-center'>{error ? error : isLoading ? "Loading..." : "CRONS"}</h1>
      {!cronList.length && <p>No CRON jobs...</p>}
      <ul>
        {!isLoading && cronList?.map((item, i) => (
          <li key={item.id + i}>
            <JobItem cronJob={item} />
          </li>
        ))
      }
      </ul>
    </div>
  )
}

export default JobList