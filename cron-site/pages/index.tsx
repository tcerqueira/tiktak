import type { NextPage } from 'next'
import Head from 'next/head'
import CronPoster, { CronFormData } from '../components/index/CronPoster'
import Header from '../components/Header'
import JobList from '../components/index/JobList'
import useCronList from '../hooks/useCronList'

const Home: NextPage = () => {
  const { cronList, isLoading, error, fetchList } = useCronList();

  return (
    <div className='bg-orange-200'>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <main className='min-h-screen pt-16 px-1 pb-[100px] max-w-[768px] mx-auto'>
        <CronPoster onPost={(d) => fetchList()}/>
        <JobList cronList={cronList} isLoading={isLoading} error={error}/>
      </main>
    </div>
  )
}

export default Home
