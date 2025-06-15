# DORA Analysis Plan: IT Separation Requirements

## 1. Objective
Identify and document requirements from EU DORA legislation regarding physical and logical separation of IT functions, applications, workloads, and tenants. Focus on separation, segregation, and isolation requirements for:
- Physical computing/hardware
- Memory space  
- Networking
- Multi-tenant environments
- Other forms of separation/segregation/isolation

## 2. Methodology

### 2.1 Source Attribution Framework
For each requirement identified, record:
1. **Requirement**: Clear statement of the MUST/SHOULD requirement
2. **Source**: Webpage title + URI
3. **Location**: Paragraph number, page number, or unique search string
4. **Context**: Sufficient detail for reader to locate precise source text

### 2.2 Analysis Approach
1. **Primary Sources First**: Start with Level 1 regulation (main legal text)
2. **Technical Standards**: Review Level 2 RTS/ITS for implementation details
3. **Guidelines**: Examine Level 3 guidelines for practical guidance
4. **Supporting Materials**: Check additional resources for clarification

### 2.3 Documentation Structure
- Create individual markdown files in `docs/notes/dora/` for each major document analyzed
- Use consistent format: `source-title-analysis.md`
- Cross-reference requirements across documents
- Maintain master requirements list in `docs/notes/dora/requirements-matrix.md`

## 3. First-Level Resources from Root URI
*Source: https://www.eiopa.europa.eu/digital-operational-resilience-act-dora_en*

### 3.1 Level 1 - Primary Legal Texts
3.1.1 **Main DORA Regulation**
   - Title: "Regulation (EU) 2022/2554 on digital operational resilience for the financial sector"
   - URI: https://eur-lex.europa.eu/eli/reg/2022/2554/oj
   - Priority: **HIGH** - Primary legal text
   - Expected content: Core separation requirements
   - Output location: `docs/notes/dora/regulation-2022-2554-analysis.md`

3.1.2 **Amending Directive**
   - Title: "Directive (EU) 2022/2556 amending various directives for digital operational resilience"
   - URI: https://eur-lex.europa.eu/eli/dir/2022/2556/oj/eng
   - Priority: **MEDIUM** - May contain sector-specific separation requirements
   - Output location: `docs/notes/dora/directive-2022-2556-analysis.md`

### 3.2 Level 2 - Technical Standards (High Priority)
3.2.1 **ICT Risk Management Framework RTS**
   - URI: https://eur-lex.europa.eu/legal-content/EN/TXT/... (RTS on ICT risk management framework)
   - Priority: **HIGH** - Likely contains technical separation requirements
   - Output location: `docs/notes/dora/rts-ict-risk-management-analysis.md`

3.2.2 **ICT Third-Party Policy RTS**
   - URI: https://eur-lex.europa.eu/legal-content/EN/TXT/... (RTS on ICT third-party policy)
   - Priority: **HIGH** - May address separation between internal/external systems
   - Output location: `docs/notes/dora/rts-ict-third-party-analysis.md`

3.2.3 **Threat Led Penetration Testing RTS**
   - URI: https://ec.europa.eu/transparency/documents-reg... (RTS on TLPT)
   - Priority: **MEDIUM** - May contain testing environment separation requirements
   - Output location: `docs/notes/dora/rts-tlpt-analysis.md`

3.2.4 **Subcontracting RTS**
   - URI: https://ec.europa.eu/transparency/documents-reg... (RTS on subcontracting)
   - Priority: **MEDIUM** - May address separation in outsourced environments
   - Output location: `docs/notes/dora/rts-subcontracting-analysis.md`

### 3.3 Level 2 - Implementation Standards
3.3.1 **ICT Incidents Classification RTS**
   - URI: https://eur-lex.europa.eu/legal-content/EN/TXT/... (RTS on ICT incidents classification)
   - Priority: **LOW** - Unlikely to contain separation requirements
   - Output location: `docs/notes/dora/rts-incidents-classification-analysis.md`

3.3.2 **ICT Incidents Reporting Process RTS**
   - URI: https://eur-lex.europa.eu/legal-content/EN/TXT/... (RTS on ICT incidents reporting process)
   - Priority: **LOW** - Unlikely to contain separation requirements
   - Output location: `docs/notes/dora/rts-incidents-reporting-analysis.md`

### 3.4 Level 3 - Guidelines
3.4.1 **Oversight Cooperation Guidelines**
   - URI: https://www.eiopa.europa.eu/publications/joint-guidelines-oversight-cooper...
   - Priority: **LOW** - Administrative focus, unlikely separation content
   - Output location: `docs/notes/dora/guidelines-oversight-cooperation-analysis.md`

### 3.5 Supporting Resources
3.5.1 **ESAs Report on ICT Third-Party Providers Landscape**
    - URI: https://www.eiopa.europa.eu/esas-publish-report...
    - Priority: **MEDIUM** - May contain separation best practices
    - Output location: `docs/notes/dora/report-ict-third-party-landscape-analysis.md`

3.5.2 **ESAs Public Statement on DORA Application**
    - URI: https://www.eiopa.europa.eu/document/download/0...
    - Priority: **MEDIUM** - May clarify separation requirements
    - Output location: `docs/notes/dora/statement-dora-application-analysis.md`

## 4. Analysis Sequence

### 4.1 Phase 1: Core Legal Requirements (Session 1)
4.1.1 **Main DORA Regulation Analysis** (EU 2022/2554)
   - Task: Download and analyze PDF
   - Task: Search for separation keywords (see section 5.1)
   - Task: Focus on Articles related to ICT risk management
   - Task: Document all separation requirements with precise citations
   - Output location: `docs/notes/dora/regulation-2022-2554-analysis.md`

### 4.2 Phase 2: Technical Implementation (Session 2)
4.2.1 **ICT Risk Management Framework RTS Analysis**
   - Task: Analyze technical requirements for separation
   - Task: Look for specific implementation guidance
   - Output location: `docs/notes/dora/rts-ict-risk-management-analysis.md`

4.2.2 **ICT Third-Party Policy RTS Analysis**
   - Task: Focus on separation between internal and third-party systems
   - Task: Document multi-tenancy requirements
   - Output location: `docs/notes/dora/rts-ict-third-party-analysis.md`

### 4.3 Phase 3: Specialized Areas (Session 3)
4.3.1 **Threat Led Penetration Testing RTS Analysis**
   - Task: Testing environment separation requirements
   - Task: Production vs. test environment isolation
   - Output location: `docs/notes/dora/rts-tlpt-analysis.md`

4.3.2 **Subcontracting RTS Analysis**
   - Task: Outsourcing separation requirements
   - Task: Multi-tenant service provider obligations
   - Output location: `docs/notes/dora/rts-subcontracting-analysis.md`

### 4.4 Phase 4: Supporting Analysis (Session 4)
4.4.1 **Supporting Reports and Statements Analysis**
   - Task: Cross-reference requirements across all documents
   - Task: Identify implementation examples
   - Task: Document industry best practices
   - Output location: `docs/notes/dora/supporting-materials-analysis.md`

4.4.2 **Master Requirements Compilation**
   - Task: Create comprehensive requirements matrix
   - Task: Cross-reference all separation requirements
   - Task: Generate implementation guidance for Milo
   - Output location: `docs/notes/dora/requirements-matrix.md`

## 5. Search Keywords and Methodology

### 5.1 Primary Keywords for Each Document
- "separat" / "separation"
- "segregat" / "segregation" 
- "isolat" / "isolation"
- "physical"
- "logical"
- "network" / "networking"
- "hardware"
- "memory"
- "tenant" / "multi-tenant" / "multi-tenancy"
- "application" / "applications"
- "workload" / "workloads"
- "environment"
- "infrastructure"
- "computing"
- "resource"

## 6. Constraints
- **Domain**: europa.eu only
- **Language**: English only
- **Hops**: Maximum 2 from root URI
- **ACU Limit**: Stop every 2 ACUs for user permission
- **Documentation**: GitHub-flavoured markdown in docs/notes/dora/

## 7. Expected Deliverables
7.1 **Individual analysis files** for each major document (see output locations in sections 3.1-3.5)
7.2 **Master requirements matrix** with all separation requirements (`docs/notes/dora/requirements-matrix.md`)
7.3 **Source attribution table** with precise citations (embedded in each analysis file)
7.4 **Implementation guidance summary** for Milo task driver plugin (`docs/notes/dora/milo-implementation-guidance.md`)

## 8. Next Steps
8.1 Begin with Phase 1 - analyze main DORA regulation (task 4.1.1)
8.2 Create `docs/notes/dora/regulation-2022-2554-analysis.md`
8.3 Extract and document all separation requirements with precise citations
8.4 Seek user permission before proceeding to Phase 2

---
*Plan created: June 15, 2025*
*Root URI: https://www.eiopa.europa.eu/digital-operational-resilience-act-dora_en*
