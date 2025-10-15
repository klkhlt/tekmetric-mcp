# Example Query Verification Checklist

This document verifies which examples in the README are actually possible with current API capabilities.

## Legend
- ✅ **Directly Supported** - Can do with current API search/filter
- ⚠️ **Requires Claude Analysis** - Data is available, but Claude needs to analyze/filter
- ❌ **Not Possible** - API doesn't support this search/filter
- 🔄 **Needs Multiple Calls** - Requires combining data from multiple tools

---

## "What You Can Do" Section

### Ask About Your Shop

- [+] ✅ `Show me today's appointments`
  - **API Support:** ✅ appointments tool with start_date/end_date filters
  - **Works:** Yes, direct date filtering

- [+] ✅ `What repair orders are in progress?`
  - **API Support:** ✅ repair_orders tool with status filter
  - **Works:** Yes, filter by status="wip"

- [-] ⚠️ `How many customers do I have?` (sort of works)
  - **API Support:** ⚠️ customers tool returns paginated list
  - **Works:** Claude can count from totalElements in response

### Find Customers & Vehicles

- [+] ✅ `Find customers named John Smith`
  - **API Support:** ✅ customers tool with query parameter
  - **Works:** Yes, searches name/email/phone

- [+] ✅ `Search for a 2020 Honda Accord`
  - **API Support:** ✅ vehicles tool with query parameter
  - **Works:** Yes, searches make/model/year/VIN

- [+] ✅ `Look up vehicle with VIN 1HGBH41JXMN109186`
  - **API Support:** ✅ vehicles tool with query parameter
  - **Works:** Yes, direct VIN search

### Track Work & Revenue

- [+] ✅ `Show me all estimates from last week`
  - **API Support:** ✅ repair_orders with status + date filters
  - **Works:** Yes, status="estimate" + date range

- [ ] ⚠️ `What's our total revenue for the month?`
  - **API Support:** ⚠️ repair_orders with date filter, Claude sums totals
  - **Works:** Claude can sum TotalSales from results
  - **Note:** Limited to 250 results, might need weekly chunks

- [ ] ⚠️ `Find repair orders over $1000`
  - **API Support:** ❌ No price filter in API
  - **Works:** Claude must filter results client-side
  - **Note:** Get all ROs, Claude filters by totalSales > 100000 (cents)

### Manage Your Schedule

- [ ] ✅ `Who's scheduled for tomorrow?`
  - **API Support:** ✅ appointments with start_date/end_date
  - **Works:** Yes, filter by tomorrow's date

- [ ] ✅ `What appointments does customer 456 have?`
  - **API Support:** ✅ appointments with customer_id filter
  - **Works:** Yes, direct customer filter

- [ ] ✅ `Show me this week's workload`
  - **API Support:** ✅ appointments with date range
  - **Works:** Yes, start_date/end_date for week

### Analyze Your Business

- [ ] 🔄 `Which customers haven't been in for 6 months?`
  - **API Support:** ❌ No "last visit date" filter
  - **Works:** Needs workaround:
    1. Get all customers
    2. For each, check recent repair orders
    3. Claude identifies customers without recent ROs
  - **Note:** **REMOVE THIS** - Too complex, requires too many API calls

- [ ] ⚠️ `What's the average repair order value this quarter?`
  - **API Support:** ⚠️ Get ROs for date range, Claude calculates average
  - **Works:** Yes, but limited to 250 results
  - **Note:** Should suggest breaking into monthly chunks

- [ ] 🔄 `Show me the most common services we perform`
  - **API Support:** 🔄 Get repair orders → analyze jobs
  - **Works:** Partially - limited by 250 RO limit
  - **Note:** **SIMPLIFY** to "Show me services from this month"

---

## "Example Queries" Section

### Daily Operations

- [ ] ✅ `What's on the schedule today?`
  - **API Support:** ✅ Same as "Show me today's appointments"
  - **Works:** Yes

- [ ] ✅ `Show me repair orders that need to be completed this week`
  - **API Support:** ⚠️ No "due date" - can filter by status + created date
  - **Works:** Partial - can show WIP orders from this week
  - **Note:** **REPHRASE** to "Show me repair orders in progress"

- [ ] ✅ `Find the customer with phone 555-1234`
  - **API Support:** ✅ customers query searches phone numbers
  - **Works:** Yes

### Business Intelligence

- [ ] ⚠️ `How many new customers did we get last month?`
  - **API Support:** ❌ No "created date" filter on customers
  - **Works:** **REMOVE** - Not directly possible
  - **Alternative:** "Show me customers" (Claude can check createdDate)

- [ ] ⚠️ `What's our average repair order value?`
  - **API Support:** ⚠️ Same as above
  - **Works:** Yes with date range

- [ ] 🔄 `Which services generate the most revenue?`
  - **API Support:** 🔄 Get ROs → analyze jobs
  - **Works:** Limited
  - **Note:** **SIMPLIFY** to "Show me this month's jobs"

### Customer Service

- [ ] 🔄 `When was this customer last here?`
  - **API Support:** 🔄 Get ROs for customer, Claude finds most recent
  - **Works:** Yes
  - **Note:** **REPHRASE** to "Show me repair orders for customer John Smith"

- [ ] ✅ `What work have we done on this vehicle?`
  - **API Support:** ✅ repair_orders with vehicle_id
  - **Works:** Yes

- [ ] 🔄 `Show me all open estimates for customer John Smith`
  - **API Support:** 🔄 Search customer → get customer_id → filter ROs
  - **Works:** Yes but multi-step
  - **Note:** Works fine, Claude can handle this

---

## Recommendations

### ❌ Remove (Too Complex or Not Possible)
1. "Which customers haven't been in for 6 months?" - Requires too many API calls
2. "Show me the most common services we perform" - Too complex for 250 limit
3. "How many new customers did we get last month?" - No customer creation date filter
4. "Which services generate the most revenue?" - Too complex

### ✏️ Rephrase (Make More Direct)
1. "Show me repair orders that need to be completed this week" → "Show me repair orders in progress"
2. "When was this customer last here?" → "Show me repair orders for customer John Smith"
3. "What's the average repair order value this quarter?" → "What's the average repair order value for last month?" (more realistic scope)

### ✅ Keep As-Is (Work Great)
- All appointment queries
- Direct customer/vehicle searches
- Status-based RO filters
- Date-range queries
- VIN/phone/name lookups

---

## Summary Statistics

- ✅ **Directly Supported:** 12 examples
- ⚠️ **Requires Claude Analysis:** 4 examples
- 🔄 **Multiple Calls Needed:** 5 examples
- ❌ **Should Remove:** 4 examples

**Total Examples:** 21
**Solid Examples:** 16 (76%)
**Need Changes:** 5 (24%)
