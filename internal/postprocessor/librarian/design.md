# Go LIbrarian Migration

# Objective

*In 1-2 sentences, describe the business goal of this project. At a high level, what should be different once this design (or an alternative) is in place?*

*For example: "gShoe should keep the user's feet warm and dry even when they walk through the snow."*

*This section should **not**:*

* *Describe the problem (use "Background")*  
* *Propose a solution (use "Design")*  
* *Be a detailed list of requirements (use [go/appeco:new-prd](https://goto.google.com/appeco:new-prd) to create a PRD or link an existing one above).* 

# Background

*Provide context for an unfamiliar reader to understand the proposal.*

* *What is the problem?*  
* *Why is it important to solve?*  
* *What solutions do we already have that solve similar problems?*

*This section should **not**:*

* *Be a lengthy explanation with a lot of details. Instead, prefer to link to other docs where available.*  
* *Include information about your design*  
* *Talk about ideas to solve the problem.*

# Overview

*High-level overview, ideally no longer than half a page; put details in the next section and background in the previous section. Should be understandable by a new Google engineer not working on the project. Diagrams can be especially useful to quickly convey the shape of the solution. This section does not need to prove that the solution meets the objective or why it is better than alternatives.*

# Detailed Design

*If you are designing a frontend system, you can view a description of this section [here](https://engdoc.corp.google.com/eng/doc/design_doc_templates/designdoc-template-frontend.html?cl=head#detailed_design). If you are designing a backend system, you can view a description of this section [here](https://engdoc.corp.google.com/eng/doc/design_doc_templates/designdoc-template-backend.html?cl=head#detailed_design). If UX is required, you should include a link to [go/appeco:new-ux-plan](https://goto.google.com/appeco:new-ux-plan). You might want to embed code via [code blocks](https://goto.google.com/smartchips#zippy=%2Cinsert-a-code-block), like this:*

```
// Example code can go here.
```

*Weigh the pros and cons of the approach you are recommending and provide a justification for why you chose this approach over those discussed in the “Alternatives considered” section below. For smaller details discuss alternatives inline while reserving the “Alternatives considered” section for overall alternative designs.*

*Make sure readers can find your code.  The tracking issue / hotlist listed in the header is fine if CLs are linked to it; if not, list the directories where the code will go.*

*Your design should include details on how you will validate and test your proposal, especially if there are specific risk(s) associated with your technical approach.*

# Alternatives considered

*Clearly list the other potential approaches to meeting the objective that you considered and why the current proposal was ultimately selected. In the rare cases where requirements or system constraints only allow for one possible high-level approach, that should be highlighted here and alternatives to specific details should still be discussed in-line in the detailed design section.*

# Work Estimates

*In addition to measuring a project's quality, you need to measure its progress. Estimate how long each phase will take (please be detailed; subtask granularity should be roughly one week). Use [go/appeco:new-execution-plan](https://goto.google.com/appeco:new-execution-plan) if you need a detailed template.* 

# Documentation plan

*What documentation will be provided (i.e., add new docs or updated existing docs) to the users and developers of this product?  When will this documentation be available to them?  If contextual documentation is required within this product, please state why.*

*The documentation can be made available to the users and developers closer to or immediately after the completion of the effort.  After creating the documentation, update the plan in the design doc with the date of publication of documentation along with links to the documentation.*

# Launch plans

*What are the launch plans for your project? This includes, but is not limited to:*

* *What visible changes* will *your project cause on the site*?  
* *What will be the impact on production and/or partners*?  
* *What new servers will be introduced*?  
* *What kind of supportability will be needed long-term (i.e. UX changes, deprecations, and migrations)*  
* *Rough timeline for releasing your project in different languages and countries.*  
* *Provide a link to [go/appeco:new-ug](https://goto.google.com/appeco:new-ug) if relevant. See also [Production Launch Review](http://go/plr-start).*

# Risks

*Describe any known risks to project implementation or launch timeline and plans to mitigate or address them (pre or post launch). These can include timeline pressure, external dependencies (operational, organizational and technical), technical debt, incomplete investigations/explorations, potential performance impacts, infrastructure related concerns etc.*