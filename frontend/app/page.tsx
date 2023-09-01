'use client';

import React from "react";
import {useRouter} from 'next/navigation'

export default function Home() {
  const router = useRouter()
  const [roomName, setroomName] = React.useState('hogehoge');

  const onChange = (roomName: string) => {
    setroomName(roomName);
  }
  const goTo = () => {
    router.push(`/${roomName}`)
  }
  return (
    <div className="relative overflow-hidden">
      <div className="max-w-[85rem] mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-24">
        <div className="text-center">
          <h1 className="text-4xl sm:text-6xl font-bold text-gray-800 dark:text-gray-200" >
            Insights
          </h1>

          <p className="mt-3 text-gray-600 dark:text-gray-400">
            Stay in the know with insights from industry experts.
          </p>

          <div className="mt-7 sm:mt-12 mx-auto max-w-xl relative">

            <div className="flex flex-row">
              <div className="basis-3/4">
                <input onChange={(e) => onChange(e.target.value)} value={roomName} type="email" id="email" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@flowbite.com" />
              </div>
              <div className="basis-1/4">
                <button onClick={goTo} type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                  Go
                </button>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div >
  )
}
