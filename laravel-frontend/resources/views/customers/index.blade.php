<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Customer Management</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 900px; margin: 40px auto; padding: 0 20px; }
        table { width: 100%; border-collapse: collapse; margin-top: 20px; }
        th, td { border: 1px solid #ddd; padding: 10px; text-align: left; }
        th { background-color: #f4f4f4; }
        .btn { display: inline-block; padding: 6px 12px; border-radius: 4px; text-decoration: none; color: white; font-size: 14px; }
        .btn-primary { background-color: #3b82f6; }
        .btn-edit { background-color: #f59e0b; }
        .btn-delete { background-color: #ef4444; border: none; cursor: pointer; }
        .alert { padding: 10px; border-radius: 4px; margin-bottom: 15px; }
        .alert-success { background-color: #d1fae5; color: #065f46; }
        .alert-error { background-color: #fee2e2; color: #991b1b; }
        .family-list { font-size: 13px; color: #555; }
    </style>
</head>
<body>
    <h1>Customer Management</h1>

    @if (session('success'))
        <div class="alert alert-success">{{ session('success') }}</div>
    @endif

    @if ($errors->any())
        <div class="alert alert-error">{{ $errors->first() }}</div>
    @endif

    <a href="{{ route('customers.create') }}" class="btn btn-primary">+ Add New Customer</a>

    <table>
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Nationality</th>
                <th>Phone</th>
                <th>Email</th>
                <th>Family Members</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($customers as $customer)
                <tr>
                    <td>{{ $customer['cst_id'] }}</td>
                    <td>{{ trim($customer['cst_name']) }}</td>
                    <td>{{ $customer['nationality']['nationality_name'] ?? '-' }}</td>
                    <td>{{ $customer['cst_phoneNum'] }}</td>
                    <td>{{ $customer['cst_email'] }}</td>
                    <td class="family-list">
                        @forelse ($customer['family'] ?? [] as $member)
                            {{ $member['fl_relation'] }}: {{ $member['fl_name'] }}<br>
                        @empty
                            <em>None</em>
                        @endforelse
                    </td>
                    <td>
                        <a href="{{ route('customers.edit', $customer['cst_id']) }}" class="btn btn-edit">Edit</a>
                        <form action="{{ route('customers.destroy', $customer['cst_id']) }}" method="POST" style="display:inline;" onsubmit="return confirm('Delete this customer?');">
                            @csrf
                            @method('DELETE')
                            <button type="submit" class="btn btn-delete">Delete</button>
                        </form>
                    </td>
                </tr>
            @empty
                <tr>
                    <td colspan="7">No customers found.</td>
                </tr>
            @endforelse
        </tbody>
    </table>
</body>
</html>