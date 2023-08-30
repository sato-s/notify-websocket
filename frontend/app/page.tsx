export default function Home() {
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
            <form>
              <input type="email" id="email" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@flowbite.com" />
            </form>
          </div>

        </div>
      </div>
    </div>
  )
}
