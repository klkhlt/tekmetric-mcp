---
sidebar_position: 6
title: Contributing
---

# Contributing to Tekmetric MCP Server

Welcome! This is an **AI-first project**, which means we build features primarily using AI coding assistants rather than hand-written code.

:::info AI-First Development
This project is built with AI assistance (primarily Claude). We encourage contributors to use AI tools for development and to contribute by providing clear requirements rather than code.
:::

## How to Contribute

There are several ways you can help improve this project:

### 1. Report Issues 🐛

Found a bug? Have a feature request? [Open an issue on GitHub](https://github.com/beetlebugorg/tekmetric-mcp/issues).

**Good issue reports include:**
- What you expected to happen
- What actually happened
- Steps to reproduce the problem
- Your environment (OS, AI client, Tekmetric environment)

### 2. Request Features ✨

We welcome feature requests! The best way to request a feature is to provide a clear requirements document.

### 3. Write Requirements Documents 📝

**This is the preferred way to contribute new features!**

Instead of submitting code, write a clear requirements document describing what you want. AI assistants (like Claude) can then use your requirements to generate high-quality code.

## How to Write a Good Requirements Document

A good requirements document helps AI understand exactly what you want built. Here's how to write one:

### Template

```markdown
# Feature Request: [Feature Name]

## Problem Statement
Describe the problem this feature solves.
- Who experiences this problem?
- When does it occur?
- Why is it a problem?

## Proposed Solution
Describe what you want the feature to do.
- What should it accomplish?
- How should it behave?
- What are the expected inputs and outputs?

## User Stories
Describe how users will interact with this feature.

Example:
- As a shop owner, I want to [action] so that [benefit]
- As a service advisor, I want to [action] so that [benefit]

## Technical Requirements
If you have specific technical requirements, list them here.
- API endpoints needed
- Data structures
- Performance requirements
- Security considerations

## Examples
Provide concrete examples of how this would work.

### Example 1: [Scenario Name]
**Input:** [What the user does]
**Expected Output:** [What should happen]

### Example 2: [Another Scenario]
**Input:** [What the user does]
**Expected Output:** [What should happen]

## Edge Cases
What unusual situations should be handled?
- Error conditions
- Empty data
- Invalid inputs
- Rate limits

## Success Criteria
How will we know this feature is complete and working correctly?
- [ ] Criterion 1
- [ ] Criterion 2
- [ ] Criterion 3

## Additional Context
Any other information that would be helpful.
- Screenshots
- Links to related features
- Tekmetric API documentation references
```

### Good Example

```markdown
# Feature Request: Search Repair Orders by Date Range

## Problem Statement
Currently, users can't easily find repair orders from a specific time period. Shop owners need to analyze business trends by looking at ROs from specific date ranges (e.g., "last month", "Q1 2024").

## Proposed Solution
Add date range filtering to the `repair_orders` tool that accepts:
- Start date (optional)
- End date (optional)
- Natural language date ranges like "last month", "this week"

## User Stories
- As a shop owner, I want to see all repair orders from last month so that I can calculate monthly revenue
- As a manager, I want to compare this quarter's ROs to last quarter's so I can track growth
- As an advisor, I want to see today's completed ROs so I can follow up with customers

## Technical Requirements
- Use Tekmetric's API date filtering (if available)
- Support ISO date format (YYYY-MM-DD)
- Handle timezone conversion properly
- Should work with existing `status` and `customer_id` filters

## Examples

### Example 1: Specific Date Range
**Input:** "Show me repair orders from January 1 to January 31, 2024"
**Expected Output:** List of all ROs created between those dates

### Example 2: Relative Date
**Input:** "Find repair orders from last week"
**Expected Output:** List of ROs from 7 days ago to today

### Example 3: Combined Filters
**Input:** "Show me completed repair orders for customer 123 from this month"
**Expected Output:** ROs matching customer ID 123, status complete, created this month

## Edge Cases
- What if start_date is after end_date? (Return error)
- What if no ROs exist in date range? (Return empty list with message)
- What about timezone differences? (Use shop's configured timezone)
- Invalid date format? (Return helpful error message)

## Success Criteria
- [ ] Can filter by exact date range
- [ ] Can filter by start date only (all ROs after date)
- [ ] Can filter by end date only (all ROs before date)
- [ ] Works combined with other filters
- [ ] Clear error messages for invalid dates
- [ ] Documentation updated with examples
```

### Bad Example

```markdown
# Add date filtering

We need date filtering for repair orders.
```

**Why this is bad:**
- Too vague
- No examples
- No success criteria
- Doesn't explain the problem or use cases

## Development Workflow

Once you've written a requirements document:

1. **Open an issue** on GitHub with your requirements document
2. **Discussion** - Maintainers and community will discuss the requirements
3. **Refinement** - Requirements may be clarified or adjusted
4. **Implementation** - Someone (often using an AI assistant) will implement the feature
5. **Review** - Pull request will be reviewed and merged
6. **Documentation** - Feature will be documented

## Code Contributions

If you prefer to submit code directly, that's fine too! We still use AI assistants to review and improve code submissions.

**When submitting code:**
- Write clear commit messages
- Include tests if applicable
- Update documentation
- Follow existing code patterns
- Explain your changes in the PR description

## AI-First Principles

This project follows these principles:

### 1. **Requirements Over Code**
Clear requirements are more valuable than code. Good requirements can generate great code multiple times.

### 2. **Documentation First**
Document what you want before building it. Documentation becomes the requirements.

### 3. **Iterate with AI**
Use AI to rapidly prototype, test, and refine features based on requirements.

### 4. **Human Review**
All AI-generated code is reviewed by humans for correctness and security.

### 5. **Test Everything**
Even (especially!) AI-generated code needs thorough testing.

## Questions?

- **Not sure if your idea is good?** Open a discussion on GitHub
- **Need help writing requirements?** Ask! We're happy to help
- **Want to pair program with AI?** Share your approach in the discussions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## Code of Conduct

Be respectful, helpful, and constructive. We're all here to build something useful together.

---

**Thank you for contributing!** Whether you write requirements, report bugs, or submit code, every contribution makes this project better. 🎉
