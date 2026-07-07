<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ $customer ? 'Edit Customer' : 'Add Customer' }}</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 700px; margin: 40px auto; padding: 0 20px; }
        label { display: block; margin-top: 15px; font-weight: bold; }
        input, select { width: 100%; padding: 10px; margin-top: 5px; box-sizing: border-box; font-size: 15px; }
        select#nationality_id { padding: 14px 10px; font-size: 16px; height: 48px; }
        .family-row { display: flex; gap: 10px; margin-top: 10px; align-items: center; }
        .family-row input { flex: 1; }
        .btn { padding: 8px 16px; border-radius: 4px; border: none; cursor: pointer; color: white; margin-top: 15px; }
        .btn-primary { background-color: #3b82f6; }
        .btn-add { background-color: #10b981; font-size: 13px; padding: 6px 10px; }
        .btn-remove { background-color: #ef4444; font-size: 13px; padding: 6px 10px; }
        .alert-error { background-color: #fee2e2; color: #991b1b; padding: 12px; border-radius: 4px; }
        .alert-error ul { margin: 5px 0 0 20px; padding: 0; }
        .field-error { color: #ef4444; font-size: 13px; margin-top: 4px; }
        a { display: inline-block; margin-top: 20px; }
    </style>
</head>
<body>
    <h1>{{ $customer ? 'Edit Customer' : 'Add New Customer' }}</h1>

    @if ($errors->any())
        <div class="alert-error">
            <strong>Please fix the following:</strong>
            <ul>
                @foreach ($errors->all() as $error)
                    <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
    @endif

    <form action="{{ $action }}" method="POST">
        @csrf
        @if ($method === 'PUT')
            @method('PUT')
        @endif

        <label for="cst_name">Nama</label>
        <input type="text" name="cst_name" id="cst_name" value="{{ old('cst_name', $customer['cst_name'] ?? '') }}">
        @error('cst_name') <div class="field-error">{{ $message }}</div> @enderror

        <label for="cst_dob">Tanggal Lahir</label>
        <input type="date" name="cst_dob" id="cst_dob"
               value="{{ old('cst_dob', $customer ? substr($customer['cst_dob'], 0, 10) : '') }}">
        @error('cst_dob') <div class="field-error">{{ $message }}</div> @enderror

        <label for="nationality_id">Kewarganegaraan</label>
        <select name="nationality_id" id="nationality_id">
            <option value="" disabled {{ old('nationality_id', $customer['nationality_id'] ?? '') === '' ? 'selected' : '' }}>
                Pilih kewarganegaraan
            </option>
            @foreach ($nationalities as $nationality)
                <option value="{{ $nationality['nationality_id'] }}"
                    {{ (string) old('nationality_id', $customer['nationality_id'] ?? '') === (string) $nationality['nationality_id'] ? 'selected' : '' }}>
                    {{ $nationality['nationality_name'] }}
                </option>
            @endforeach
        </select>
        @error('nationality_id') <div class="field-error">{{ $message }}</div> @enderror

        <label for="cst_phoneNum">Phone Number</label>
        <input type="text" name="cst_phoneNum" id="cst_phoneNum" value="{{ old('cst_phoneNum', $customer['cst_phoneNum'] ?? '') }}">
        @error('cst_phoneNum') <div class="field-error">{{ $message }}</div> @enderror

        <label for="cst_email">Email</label>
        <input type="email" name="cst_email" id="cst_email" value="{{ old('cst_email', $customer['cst_email'] ?? '') }}">
        @error('cst_email') <div class="field-error">{{ $message }}</div> @enderror

        <h3>Keluarga <button type="button" class="btn btn-add" onclick="addFamilyRow()">+ Tambah Keluarga</button></h3>

        <div id="family-container">
            @php $existingFamily = $customer['family'] ?? []; @endphp

            @if (count($existingFamily) > 0)
                @foreach ($existingFamily as $member)
                    <div class="family-row">
                        <input type="text" name="family_relation[]" placeholder="Relation (e.g. Father)" value="{{ $member['fl_relation'] }}">
                        <input type="text" name="family_name[]" placeholder="Nama" value="{{ $member['fl_name'] }}">
                        <input type="date" name="family_dob[]" value="{{ $member['fl_dob'] }}">
                        <button type="button" class="btn btn-remove" onclick="removeFamilyRow(this)">Hapus</button>
                    </div>
                @endforeach
            @else
                <div class="family-row">
                    <input type="text" name="family_relation[]" placeholder="Relation (e.g. Father)">
                    <input type="text" name="family_name[]" placeholder="Nama">
                    <input type="date" name="family_dob[]">
                    <button type="button" class="btn btn-remove" onclick="removeFamilyRow(this)">Hapus</button>
                </div>
            @endif
        </div>

        <br>
        <button type="submit" class="btn btn-primary">Save</button>
    </form>

    <a href="{{ route('customers.index') }}">← Back to list</a>

    <script>
        function addFamilyRow() {
            const container = document.getElementById('family-container');
            const row = document.createElement('div');
            row.className = 'family-row';
            row.innerHTML = `
                <input type="text" name="family_relation[]" placeholder="Relation (e.g. Father)">
                <input type="text" name="family_name[]" placeholder="Nama">
                <input type="date" name="family_dob[]">
                <button type="button" class="btn btn-remove" onclick="removeFamilyRow(this)">Hapus</button>
            `;
            container.appendChild(row);
        }

        function removeFamilyRow(button) {
            const container = document.getElementById('family-container');
            const row = button.closest('.family-row');
            const inputs = row.querySelectorAll('input');

            let hasContent = false;
            inputs.forEach(input => {
                if (input.value.trim() !== '') {
                    hasContent = true;
                }
            });

            if (container.children.length <= 1) {
                alert('At least one family row must remain. Clear the fields instead if not needed.');
                return;
            }

            if (hasContent) {
                if (confirm('This row has data. Are you sure you want to delete it?')) {
                    row.remove();
                }
            } else {
                row.remove();
            }
        }
    </script>
</body>
</html>