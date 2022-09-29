import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import CronPoster from '../components/CronPoster'
import Header from '../components/Header'

const Home: NextPage = () => {
  return (
    <div className='bg-orange-200'>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <main className='min-h-screen pt-16 px-1'>
        <CronPoster />
        Hello world!
      </main>
    </div>
  )
}

export default Home
