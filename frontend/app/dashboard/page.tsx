export default function DashboardPage() {
  return (
    <div className="flex min-h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-900 text-white flex-col hidden md:flex">
        <div className="p-6 text-2xl font-bold tracking-wide border-b border-gray-800">
          MyApp
        </div>
        <nav className="flex-1 p-4 space-y-2">
          <a href="#" className="flex items-center px-4 py-3 bg-gray-800 rounded-lg text-white">
            <span className="mr-3">🏠</span> Dashboard
          </a>
          <a href="#" className="flex items-center px-4 py-3 hover:bg-gray-800 rounded-lg text-gray-400 hover:text-white transition-colors">
            <span className="mr-3">👤</span> Profile
          </a>
          <a href="#" className="flex items-center px-4 py-3 hover:bg-gray-800 rounded-lg text-gray-400 hover:text-white transition-colors">
            <span className="mr-3">⚙️</span> Settings
          </a>
        </nav>
        <div className="p-4 border-t border-gray-800">
          <a href="/" className="flex items-center px-4 py-2 hover:bg-red-600 rounded-lg text-gray-400 hover:text-white transition-colors">
            <span className="mr-3">🚪</span> Logout
          </a>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 flex flex-col bg-gray-50">
        {/* Header */}
        <header className="bg-white shadow-sm px-8 py-4 flex justify-between items-center sticky top-0 z-10">
          <h1 className="text-2xl font-bold text-gray-800">Dashboard</h1>
          <div className="flex items-center space-x-4">
            <div className="w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center text-gray-500">
              U
            </div>
          </div>
        </header>

        {/* Content Body */}
        <div className="p-8 space-y-8">
          {/* Welcome Card */}
          <div className="bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl p-8 text-white shadow-lg">
            <h2 className="text-3xl font-bold mb-2">Welcome Back!</h2>
            <p className="opacity-90 max-w-2xl">
              This is your protected dashboard explicitly connected to your Go backend. 
              Your session is secured via PostgreSQL.
            </p>
          </div>

          {/* Stats Grid */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
              <div className="text-gray-500 text-sm font-medium uppercase tracking-wider">Total Users</div>
              <div className="text-3xl font-bold text-gray-900 mt-2">1,245</div>
            </div>
            <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
              <div className="text-gray-500 text-sm font-medium uppercase tracking-wider">Active Sessions</div>
              <div className="text-3xl font-bold text-gray-900 mt-2">432</div>
            </div>
            <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
              <div className="text-gray-500 text-sm font-medium uppercase tracking-wider">Server Status</div>
              <div className="text-3xl font-bold text-green-600 mt-2">Online</div>
            </div>
          </div>

          {/* Info Section */}
          <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
             <div className="px-6 py-4 border-b border-gray-100 bg-gray-50">
               <h3 className="font-semibold text-gray-700">System Information</h3>
             </div>
             <div className="p-6">
               <p className="text-gray-600 leading-relaxed">
                 You are currently authenticated. To populate this dashboard with real user-specific data, 
                 implement a protected endpoint in your Go backend (e.g., <code className="bg-gray-100 px-1 py-0.5 rounded text-pink-600 text-sm">/api/user/stats</code>) 
                 and fetch it here using the stored token.
               </p>
             </div>
          </div>
        </div>
      </main>
    </div>
  );
}
