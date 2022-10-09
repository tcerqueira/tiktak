import React from 'react'
import useCronList from '../../hooks/useCronList';
import JobItem from './JobItem'

function JobList() {
  const { cronList, isLoading, error } = useCronList();

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