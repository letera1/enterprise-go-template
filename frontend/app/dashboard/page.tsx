'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Dashboard() {
  const [user, setUser] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    fetch('http://127.0.0.1:9000/api/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Important to send cookies
    })
      .then((res) => {
        if (res.status === 401) {
          router.push('/');
          return null;
        }
        return res.json();
      })
      .then((data) => {
        if (data) {
          setUser(data);
        }
        setLoading(false);
      })
      .catch((err) => {
        console.error(err);
        setLoading(false);
      });
  }, [router]);

  const handleLogout = () => {
    fetch('http://127.0.0.1:9000/auth/logout', { 
        method: 'POST',
        credentials: 'include'
    }).then(() => {
        router.push('/');
    });
  };

  if (loading) return <div className="text-xl">Loading...</div>;

  return (
    <div className="bg-white p-8 rounded-xl shadow-lg w-full max-w-2xl">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Dashboard</h1>
        <button 
            onClick={handleLogout}
            className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition"
        >
            Logout
        </button>
      </div>
      
      <div className="space-y-4">
        <div className="p-4 bg-green-50 rounded border border-green-200">
            <h2 className="text-xl font-semibold text-green-800 mb-2">Welcome Back!</h2>
            <p className="text-green-700">You have successfully logged in.</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-6">
            <div className="p-4 bg-gray-50 rounded border">
                <p className="text-sm text-gray-500 font-semibold uppercase">ID</p>
                <p className="font-mono text-gray-800">{user?.user_id}</p>
            </div>
            <div className="p-4 bg-gray-50 rounded border">
                <p className="text-sm text-gray-500 font-semibold uppercase">Email</p>
                <p className="font-mono text-gray-800">{user?.email}</p>
            </div>
        </div>
      </div>
    </div>
  );
}
