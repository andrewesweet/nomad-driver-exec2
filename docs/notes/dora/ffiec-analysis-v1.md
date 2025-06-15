# FFIEC IT Examination Handbook Analysis: IT Separation and Threat Actor Requirements

## Document Information
- **Title**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet
- **Source URI**: https://ithandbook.ffiec.gov/it-booklets/information-security/
- **Analysis Date**: June 15, 2025
- **Document Status**: Current FFIEC guidance for financial institution examination
- **Regulatory Authority**: Federal Financial Institutions Examination Council (FFIEC)

## Analysis Methodology
Systematic search for separation, segregation, isolation, and threat actor requirements using keywords from meta-regulatory analysis framework v1.3:

**Standard Keywords**:
- Core separation terms: "separat", "segregat", "isolat"
- Physical/logical: "physical", "logical", "network", "hardware", "memory"
- Multi-tenancy: "tenant", "multi-tenant", "multi-tenancy", "shared"
- Applications: "application", "workload", "service", "function"
- Environment: "environment", "production", "testing", "development"
- Infrastructure: "infrastructure", "computing", "resource", "system"

**Threat Actor Keywords**:
- Threat actors: "threat actor", "threat actors", "adversary", "adversaries", "attacker", "attackers"
- Nation-state: "nation state", "nation-state", "state-sponsored", "foreign intelligence", "espionage"
- Advanced threats: "advanced persistent threat", "APT", "sophisticated attack"
- Criminal actors: "cybercriminal", "cyber criminal", "organized crime", "criminal organization"
- Insider threats: "insider threat", "malicious insider", "insider attack", "rogue employee"
- Other actors: "hacktivist", "hacktivism", "terrorist", "terrorism", "extremist"
- Risk analysis: "threat model", "threat modeling", "risk assessment", "threat landscape"

## Separation Requirements Identified

### 1. User Security Controls - Segregation of Duties

#### 1.1 Segregation of Duties for Critical Tasks
**Requirement**: Financial institutions must implement segregation of duties or job designs that require more than one person to complete critical or sensitive tasks to help mitigate risk.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - Opening paragraph
**Context**: "Segregation of duties, or job designs that require more than one person to complete critical or sensitive tasks, can help mitigate risk."

#### 1.2 System Administrator Privilege Separation
**Requirement**: Financial institutions must evaluate the process for determining which individuals should be granted system administrator privileges and appropriately monitor such access for unauthorized or inappropriate activity.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - First paragraph
**Context**: "System administrators, for instance, have the most powerful role in the user access process and have unlimited access to an institution's information assets and technology. Given this extensive access, management should evaluate the process for determining which individuals should be granted system administrator privileges. Such access should be appropriately monitored for unauthorized or inappropriate activity."

#### 1.3 Independent Reviews for Multiple Functions
**Requirement**: Financial institutions must incorporate independent reviews or approvals for individuals who perform multiple functions to minimize the potential for fraud, irregularities, and errors.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - First paragraph
**Context**: "Management should incorporate independent reviews or approvals for individuals who perform multiple functions to minimize the potential for fraud, irregularities, and errors."

#### 1.4 Independent Monitoring of Privileged Users
**Requirement**: Financial institutions must implement independent monitoring of activities performed by users with increased privileges, including system administrators and super users.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - Examples list
**Context**: Listed as first example of segregation of duties: "Independent monitoring of the activities performed by the users with increased privileges (e.g., system administrators and super users)."

#### 1.5 Distribution of System Administration Activities
**Requirement**: Financial institutions must distribute system administration activities so that no administrator can hide his or her activities or control an entire system.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - Examples list
**Context**: Listed as second example of segregation of duties: "Distribution of system administration activities so no administrator can hide his or her activities or control an entire system."

#### 1.6 Approval Level Escalation for Critical Decisions
**Requirement**: Financial institutions must implement additional levels of approval as the criticality and sensitivity of decisions increase.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - Examples list
**Context**: Listed as third example of segregation of duties: "Additional levels of approval as the criticality and sensitivity of decisions increase."

#### 1.7 Independent Review for Inadequate Segregation
**Requirement**: Financial institutions must require an independent review (e.g., audit) of any activity conducted without appropriate segregation of duties.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.7(c) Segregation of Duties
**Location**: Section II.C.7(c) - Final paragraph
**Context**: "If an activity is conducted without appropriate segregation of duties, management should require an independent review (e.g., audit) of that activity."

### 2. Third-Party Service Provider Oversight - Separation and Control Requirements

#### 2.1 Third-Party Due Diligence Separation
**Requirement**: Financial institutions must conduct appropriate due diligence in selecting and monitoring third-party service providers, ensuring they use suitable information security controls when providing services.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Main content paragraph
**Context**: "Management should conduct appropriate due diligence in selecting and monitoring third-party service providers. Management should be responsible for ensuring that such third parties use suitable information security controls when providing services to the institution."

#### 2.2 Third-Party Contractual Control Separation
**Requirement**: Financial institutions must establish contractual assurances for security responsibilities, controls, and reporting with third-party service providers.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Action Summary list
**Context**: Listed in oversight requirements: "Contractual assurances for security responsibilities, controls, and reporting."

#### 2.3 Third-Party Data Protection Separation
**Requirement**: Financial institutions must require nondisclosure agreements regarding the institution's systems and data when working with third-party service providers.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Action Summary list
**Context**: Listed in oversight requirements: "Nondisclosure agreements regarding the institution's systems and data."

#### 2.4 Independent Third-Party Security Review Separation
**Requirement**: Financial institutions must conduct independent review of third-party security through appropriate reports from audits and tests.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Action Summary list
**Context**: Listed in oversight requirements: "Independent review of the third party's security through appropriate reports from audits and tests."

#### 2.5 Third-Party Incident Response Coordination Separation
**Requirement**: Financial institutions must coordinate incident response policies and contractual notification requirements with third-party service providers.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Action Summary list
**Context**: Listed in oversight requirements: "Coordination of incident response policies and contractual notification requirements."

#### 2.6 Third-Party Customer Information Separation
**Requirement**: Financial institutions must require third-party service providers to implement appropriate measures by contract when they store, transmit, process, or dispose of customer information.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Customer information paragraph
**Context**: "If the third-party service provider stores, transmits, processes, or disposes of customer information, management should require third-party service providers by contract to implement appropriate measures designed to meet the Information Security Standards."

#### 2.7 Third-Party Access Rights and Audit Separation
**Requirement**: Financial institutions must specify in contracts that the institution or an independent auditor has access to the service provider to perform evaluations against Information Security Standards.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Contract requirements list
**Context**: Listed in contract requirements: "Specify that the institution or an independent auditor has access to the service provider to perform evaluations of the service provider's performance against the Information Security Standards."

#### 2.8 Third-Party Cyber Threat Risk Separation
**Requirement**: Financial institutions must determine whether cyber risks are identified, measured, mitigated, monitored, and reported by third parties, and incorporate assessment of third-party cyber threats into institutional risk reporting.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.C.20 Oversight of Third-Party Service Providers
**Location**: Section II.C.20 - Final paragraph
**Context**: "Additionally, as part of the oversight of third-party service providers, management should determine whether cyber risks are identified, measured, mitigated, monitored, and reported by such third parties as third-party cyber threats can have an impact on the institution. Information security reporting by the institution should incorporate an assessment of these third-party risks to facilitate a comprehensive understanding of the institution's exposure to third-party cyber threats."

## Threat Actor Requirements Identified

### 1. General Threat Definition and Sources

#### 1.1 NIST-Based Threat Definition
**Requirement**: Financial institutions must consider threats as defined by NIST - any circumstance or event with the potential to create loss, including persons with intent to harm or who unintentionally cause harm.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.A.1 Threats
**Location**: Section II.A.1 - Opening paragraph
**Context**: "According to the National Institute of Standards and Technology (NIST), a threat is any circumstance or event with the potential to create loss. A threat can be a natural occurrence, technology or physical failure, a person with intent to harm, or a person who unintentionally causes harm."

#### 1.2 Threat Intelligence Sources Requirement
**Requirement**: Financial institutions must utilize both public and private sources for threat information, including news media, government publications, information security vendors, and information-sharing organizations.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.A.1 Threats
**Location**: Section II.A.1 - First paragraph
**Context**: "Information about threats is available from public and private sources. Public sources include the news media, blogs, government publications and announcements, and various websites. Private sources include information security vendors and information-sharing organizations."

#### 1.3 Threat Modeling Requirement
**Requirement**: Financial institutions should consider using threat modeling as a structured approach to aggregate and quantify potential threats, understand their nature, frequency, and sophistication.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.A.1 Threats
**Location**: Section II.A.1 - Second paragraph
**Context**: "Threat modeling is a structured approach that enables an institution to aggregate and quantify potential threats. Institutions should consider using threat modeling to better understand the nature, frequency, and sophistication of threats; evaluate the information security risks to the institution; and apply this knowledge to the institution's information security program."

#### 1.4 Zero-Day Attack Consideration
**Requirement**: Financial institutions must acknowledge that threat modeling may not account for previously unseen attacks such as zero-day attacks, but these could have significant impacts.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section II.A.1 Threats
**Location**: Section II.A.1 - Second paragraph
**Context**: "As threats evolve rapidly, however, it is understood that modeling may not account for attacks that have not previously been seen, such as zero-day attacks, but could have significant impacts."

### 2. Specific Threat Source Categories

#### 2.1 NIST Threat Source Classification Requirement
**Requirement**: Financial institutions must consider the four specific types of threat sources as defined by NIST: hostile cyber or physical attacks, human errors, structural failures, and natural/man-made disasters.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Main content paragraph
**Context**: "NIST notes that types of threat sources include the following: Hostile cyber or physical attacks. Human errors of omission or commission. Structural failures of organization-controlled resources (e.g., hardware, software, and environmental controls). Natural and man-made disasters, accidents, and failures beyond the control of the organization."

#### 2.2 Hostile Cyber and Physical Attack Consideration
**Requirement**: Financial institutions must specifically assess and prepare for hostile cyber or physical attacks as a distinct threat source category.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - NIST threat source list
**Context**: Listed as first category in NIST threat source types: "Hostile cyber or physical attacks."

#### 2.3 Human Error Threat Assessment
**Requirement**: Financial institutions must consider human errors of omission or commission as threat sources requiring assessment and mitigation.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - NIST threat source list
**Context**: Listed as second category in NIST threat source types: "Human errors of omission or commission."

### 3. Threat Intelligence and Information Sources

#### 3.1 Government Threat Intelligence Requirement
**Requirement**: Financial institutions must utilize government sources for threat intelligence, specifically including US-CERT and similar government agencies.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Threat information sources paragraph
**Context**: "Information about threats generally comes from government (e.g., US-CERT), information-sharing organizations (e.g., FS-ISAC), industry sources, the institution, and third parties."

#### 3.2 Financial Services Information Sharing Requirement
**Requirement**: Financial institutions must utilize information-sharing organizations, specifically including FS-ISAC (Financial Services Information Sharing and Analysis Center).

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Threat information sources paragraph
**Context**: "Information about threats generally comes from government (e.g., US-CERT), information-sharing organizations (e.g., FS-ISAC), industry sources, the institution, and third parties."

#### 3.3 Third-Party Threat Intelligence Requirement
**Requirement**: Financial institutions must consider third-party threat intelligence from organizations that track and report on threats, including incident reports from multiple organizations worldwide.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Third-party information paragraph
**Context**: "Third-party information may be from organizations that specifically track and report on threats or from third-party reports of past activity. Some of those reports compile knowledge from incidents reported by many organizations worldwide."

#### 3.4 Managed Security Service Provider Intelligence
**Requirement**: Financial institutions must consider attack data from managed security service providers as part of their threat intelligence program.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Information types list
**Context**: Listed in threat assessment information sources: "Attack data from sources including FS-ISAC and managed security service providers."

### 4. Threat Assessment and Response Requirements

#### 4.1 Threat Taxonomy Implementation
**Requirement**: Financial institutions should use a threat taxonomy to reduce complexity of threat assessment and enable efficient understanding of risk mitigations.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Threat taxonomy paragraph
**Context**: "By using a threat taxonomy, the institution may greatly reduce the complexity of threat assessment and enable efficient understanding of reasonable risk mitigations."

#### 4.2 Threat Source Capability and Intent Assessment
**Requirement**: Financial institutions must assess specific factors including threat source description, operational context, capabilities and intent, and benefits/consequences from the threat-source perspective.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Threat assessment factors paragraph
**Context**: "Specific factors in the threat assessment may include a description, context for operation, capabilities and intent, and, from the threat-source perspectives, benefits and negative consequences associated with an attack."

#### 4.3 Immediate and Consequential Threat Response
**Requirement**: Financial institutions must design policies to allow immediate and consequential threats to be dealt with expeditiously, while less significant threats are addressed through broader risk management processes.

**Source**: FFIEC IT Examination Handbook InfoBase - Information Security Booklet, Section III.A Threat Identification and Assessment
**Location**: Section III.A - Response design paragraph
**Context**: "Management should design policies to allow for immediate and consequential threats to be dealt with expeditiously, while less significant threats are addressed as part of a broader risk management process."

## Summary of Key Separation Requirements
1. **Critical Task Segregation**: Multi-person requirement for critical/sensitive tasks (Section II.C.7(c))
2. **System Administrator Privilege Separation**: Evaluation and monitoring of privileged access (Section II.C.7(c))
3. **Independent Reviews**: Required for individuals with multiple functions (Section II.C.7(c))
4. **Privileged User Monitoring**: Independent monitoring of system administrators and super users (Section II.C.7(c))
5. **Administration Activity Distribution**: Prevention of single-administrator system control (Section II.C.7(c))
6. **Approval Level Escalation**: Additional approvals for critical/sensitive decisions (Section II.C.7(c))
7. **Independent Review Requirement**: Audit required when segregation is inadequate (Section II.C.7(c))
8. **Third-Party Due Diligence Separation**: Due diligence and monitoring of third-party service providers (Section II.C.20)
9. **Third-Party Contractual Control Separation**: Contractual security responsibilities and controls (Section II.C.20)
10. **Third-Party Data Protection Separation**: Nondisclosure agreements for institutional systems and data (Section II.C.20)
11. **Independent Third-Party Security Review**: Independent audits and tests of third-party security (Section II.C.20)
12. **Third-Party Incident Response Coordination**: Coordinated incident response policies and notifications (Section II.C.20)
13. **Third-Party Customer Information Separation**: Contractual measures for customer information protection (Section II.C.20)
14. **Third-Party Access Rights and Audit Separation**: Institutional audit access to third-party providers (Section II.C.20)
15. **Third-Party Cyber Threat Risk Separation**: Assessment and reporting of third-party cyber threat exposure (Section II.C.20)

## Summary of Key Threat Actor Requirements
1. **NIST Threat Definition**: Must consider persons with intent to harm and unintentional harm causers (Section II.A.1)
2. **Threat Intelligence Sources**: Must utilize public and private threat information sources (Section II.A.1)
3. **Threat Modeling**: Should implement structured threat modeling approaches (Section II.A.1)
4. **Zero-Day Awareness**: Must acknowledge limitations of threat modeling for novel attacks (Section II.A.1)
5. **Specific Threat Source Categories**: Must assess hostile cyber/physical attacks, human errors, structural failures, disasters (Section III.A)
6. **Government Intelligence**: Must utilize US-CERT and government threat intelligence sources (Section III.A)
7. **FS-ISAC Integration**: Must participate in financial services information sharing (Section III.A)
8. **Third-Party Intelligence**: Must consider threat intelligence from specialized tracking organizations (Section III.A)
9. **MSSP Attack Data**: Must utilize attack data from managed security service providers (Section III.A)
10. **Threat Taxonomy**: Should implement threat taxonomy for efficient risk assessment (Section III.A)
11. **Capability Assessment**: Must assess threat source capabilities, intent, and operational context (Section III.A)
12. **Response Prioritization**: Must design expeditious response policies for immediate/consequential threats (Section III.A)

## Implementation Guidance for Milo Task Driver Plugin

### Threat Actor Considerations
1. **Threat Modeling Integration**: Implement structured threat modeling capabilities to assess risks from various threat actors in multi-tenant environments
2. **Intelligence Integration**: Design interfaces to consume threat intelligence from both public and private sources, including US-CERT and FS-ISAC
3. **Zero-Day Preparedness**: Implement defensive measures that can respond to novel attack patterns not previously modeled
4. **Hostile Actor Assessment**: Implement capabilities to assess hostile cyber and physical attacks as distinct threat categories
5. **Human Error Mitigation**: Design controls to address human errors of omission and commission as threat sources
6. **Third-Party Threat Assessment**: Implement monitoring and assessment of third-party cyber threats that can impact multi-tenant environments

### Separation Requirements
1. **Duty Segregation**: Ensure role-based access controls prevent conflicts of interest in multi-tenant workload management
2. **Access Control**: Implement segregation of duties in administrative functions for tenant isolation
3. **Privileged User Monitoring**: Independent monitoring of system administrators and super users across tenant environments
4. **Third-Party Isolation**: Implement due diligence and contractual controls for third-party service providers in multi-tenant contexts
5. **Incident Response Coordination**: Separate incident response policies and notification requirements for different tenants
6. **Audit Access Separation**: Ensure independent audit access while maintaining tenant data separation

### Compliance Considerations
- Regular threat modeling exercises incorporating FFIEC guidance and NIST threat source categories
- Documentation of threat intelligence sources including government (US-CERT) and industry (FS-ISAC) sources
- Segregation of duties implementation in operational procedures with independent review requirements
- Third-party risk assessment and monitoring incorporating cyber threat evaluation
- Ongoing assessment of novel threat vectors and attack patterns with taxonomy-based risk mitigation

---
*Analysis completed: FFIEC trial run of meta-regulatory analysis framework*
*Total requirements identified: 15 separation requirements, 12 threat actor requirements*
*Framework validation: Successfully demonstrated threat actor and separation requirement identification methodology*
