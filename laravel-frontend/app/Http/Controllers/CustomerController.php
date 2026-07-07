<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class CustomerController extends Controller
{
    protected string $apiUrl;

    public function __construct()
    {
        $this->apiUrl = config('services.go_api.url');
    }

    public function index()
    {
        $response = Http::get("{$this->apiUrl}/customers");
        $customers = $response->successful() ? ($response->json() ?? []) : [];

        return view('customers.index', compact('customers'));
    }

    public function create()
    {
        $nationalities = $this->getNationalities();

        return view('customers.form', [
            'customer' => null,
            'nationalities' => $nationalities,
            'action' => route('customers.store'),
            'method' => 'POST',
        ]);
    }

    public function store(Request $request)
    {
        $request->validate([
            'nationality_id' => 'required|integer',
            'cst_name' => 'required|string|max:50',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:50',
            'family_relation.*' => 'nullable|string|max:50',
            'family_name.*' => 'nullable|string|max:50',
            'family_dob.*' => 'nullable|date',
        ], [
            'nationality_id.required' => 'Kewarganegaraan wajib dipilih.',
            'cst_name.required' => 'Nama wajib diisi.',
            'cst_dob.required' => 'Tanggal lahir wajib diisi.',
            'cst_phoneNum.required' => 'Nomor telepon wajib diisi.',
            'cst_email.required' => 'Email wajib diisi.',
            'cst_email.email' => 'Format email tidak valid.',
        ]);

        $payload = $this->buildPayload($request);

        $response = Http::post("{$this->apiUrl}/customers", $payload);

        if ($response->failed()) {
            return back()->withInput()->withErrors(['error' => 'Failed to create customer: ' . $response->body()]);
        }

        return redirect()->route('customers.index')->with('success', 'Customer created successfully.');
    }

    public function edit($id)
    {
        $response = Http::get("{$this->apiUrl}/customers/{$id}");

        if ($response->failed()) {
            return redirect()->route('customers.index')->withErrors(['error' => 'Customer not found.']);
        }

        $customer = $response->json();
        $nationalities = $this->getNationalities();

        return view('customers.form', [
            'customer' => $customer,
            'nationalities' => $nationalities,
            'action' => route('customers.update', $id),
            'method' => 'PUT',
        ]);
    }

    public function update(Request $request, $id)
    {
        $request->validate([
            'nationality_id' => 'required|integer',
            'cst_name' => 'required|string|max:50',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:50',
            'family_relation.*' => 'nullable|string|max:50',
            'family_name.*' => 'nullable|string|max:50',
            'family_dob.*' => 'nullable|date',
        ], [
            'nationality_id.required' => 'Kewarganegaraan wajib dipilih.',
            'cst_name.required' => 'Nama wajib diisi.',
            'cst_dob.required' => 'Tanggal lahir wajib diisi.',
            'cst_phoneNum.required' => 'Nomor telepon wajib diisi.',
            'cst_email.required' => 'Email wajib diisi.',
            'cst_email.email' => 'Format email tidak valid.',
        ]);

        $payload = $this->buildPayload($request);

        $response = Http::put("{$this->apiUrl}/customers/{$id}", $payload);

        if ($response->failed()) {
            return back()->withInput()->withErrors(['error' => 'Failed to update customer: ' . $response->body()]);
        }

        return redirect()->route('customers.index')->with('success', 'Customer updated successfully.');
    }

    public function destroy($id)
    {
        $response = Http::delete("{$this->apiUrl}/customers/{$id}");

        if ($response->failed()) {
            return back()->withErrors(['error' => 'Failed to delete customer.']);
        }

        return redirect()->route('customers.index')->with('success', 'Customer deleted successfully.');
    }

    private function getNationalities(): array
    {
        $response = Http::get("{$this->apiUrl}/nationalities");
        return $response->successful() ? ($response->json() ?? []) : [];
    }

    private function buildPayload(Request $request): array
    {
        $family = [];

        if ($request->has('family_relation')) {
            foreach ($request->family_relation as $index => $relation) {
                if (empty($relation) && empty($request->family_name[$index])) {
                    continue;
                }

                $family[] = [
                    'fl_relation' => $relation,
                    'fl_name' => $request->family_name[$index] ?? '',
                    'fl_dob' => $request->family_dob[$index] ?? '',
                ];
            }
        }

        return [
            'nationality_id' => (int) $request->nationality_id,
            'cst_name' => $request->cst_name,
            'cst_dob' => $request->cst_dob . 'T00:00:00Z',
            'cst_phoneNum' => $request->cst_phoneNum,
            'cst_email' => $request->cst_email,
            'family' => $family,
        ];
    }
}