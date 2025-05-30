import React from 'react';
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card';
import {Link, Outlet} from 'react-router-dom';
import {ClipboardList, List} from 'lucide-react';

const Admin: React.FC = () => {
    return (
        <div className="container mx-auto px-4">
            <Card className="w-full shadow-lg">
                <CardHeader className="border-b border-gray-200">
                    <CardTitle className="text-3xl font-bold text-gray-900">管理者ページ</CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="flex flex-col sm:flex-row gap-6 mb-8">
                        <Link
                            to="assign-allowance"
                            className="flex items-center gap-2 px-4 py-2 bg-blue-50 rounded hover:bg-blue-100 transition-colors"
                        >
                            <ClipboardList className="w-6 h-6 text-blue-600"/>
                            <span className="text-blue-600 font-medium">従業員への手当割り当て</span>
                        </Link>
                        <Link
                            to="clock-requests"
                            className="flex items-center gap-2 px-4 py-2 bg-blue-50 rounded hover:bg-blue-100 transition-colors"
                        >
                            <List className="w-6 h-6 text-blue-600"/>
                            <span className="text-blue-600 font-medium">打刻修正申請一覧</span>
                        </Link>
                    </div>
                    <Outlet/>
                </CardContent>
            </Card>
        </div>
    );
};

export default Admin;
