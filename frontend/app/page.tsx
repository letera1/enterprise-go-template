import Link from 'next/link';

export default function Home() {
  return (
    <div className="z-10 max-w-md w-full items-center justify-center font-mono text-sm flex flex-col gap-6 bg-white p-8 rounded-xl shadow-lg">
      <h1 className="text-3xl font-bold text-gray-800">Welcome</h1>
      <p className="text-gray-600 text-center">Please sign in to continue</p>
      
      <div className="flex flex-col gap-3 w-full">
        <a 
          href="http://127.0.0.1:9000/auth/github/login" 
          className="flex items-center justify-center gap-2 bg-[#24292e] text-white px-6 py-3 rounded-lg hover:opacity-90 transition font-medium"
        >
          <span>Login with GitHub</span>
        </a>
        
        <a 
          href="http://127.0.0.1:9000/auth/google/login" 
          className="flex items-center justify-center gap-2 bg-white border border-gray-300 text-gray-700 px-6 py-3 rounded-lg hover:bg-gray-50 transition font-medium"
        >
          <span>Login with Google</span>
        </a>
      </div>
    </div>
  );
}
